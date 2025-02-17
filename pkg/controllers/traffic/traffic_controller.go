/*
Copyright 2022 The MultiCluster Traffic Controller Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package traffic

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	"github.com/Kuadrant/multi-cluster-traffic-controller/pkg/_internal/metadata"
	"github.com/Kuadrant/multi-cluster-traffic-controller/pkg/traffic"
)

// Reconciler reconciles a traffic object
type Reconciler struct {
	WorkloadClient client.Client
	ControlClient  client.Client
}

func (r *Reconciler) Handle(ctx context.Context, o runtime.Object) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	trafficAccessor := o.(traffic.Interface)
	log.Log.Info("got traffic object", "kind", trafficAccessor.GetKind(), "name", trafficAccessor.GetName(), "namespace", trafficAccessor.GetNamespace())

	if trafficAccessor.GetAnnotations() == nil {
		return ctrl.Result{}, fmt.Errorf("expected dummy configmap annotation missing")
	}

	cmName := trafficAccessor.GetAnnotations()["configmap"]
	cmField := trafficAccessor.GetAnnotations()["field"]

	cm := &corev1.ConfigMap{}

	err := r.WorkloadClient.Get(ctx, client.ObjectKey{Namespace: trafficAccessor.GetNamespace(), Name: cmName}, cm)
	if err != nil {
		return ctrl.Result{}, err
	}

	value := cm.Data[cmField]

	metadata.AddLabel(trafficAccessor, "dummy-configmap-value", value)

	return ctrl.Result{}, nil
}
