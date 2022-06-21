//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by ctrlkit. DO NOT EDIT.

/*
 * Copyright 2022 Singularity Data
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package manager

import (
	"context"
	"fmt"

	"github.com/go-logr/logr"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	risingwavev1alpha1 "github.com/singularity-data/risingwave-operator/apis/risingwave/v1alpha1"
	"github.com/singularity-data/risingwave-operator/pkg/ctrlkit"
)

// RisingWaveControllerManagerState is the state manager of RisingWaveControllerManager.
type RisingWaveControllerManagerState struct {
	client.Reader
	target *risingwavev1alpha1.RisingWave
}

// GetComputeStatefulSet gets computeStatefulSet with name equals to ${target.Name}-compute.
func (s *RisingWaveControllerManagerState) GetComputeStatefulSet(ctx context.Context) (*appsv1.StatefulSet, error) {
	var computeStatefulSet appsv1.StatefulSet

	err := s.Get(ctx, types.NamespacedName{
		Namespace: s.target.Namespace,
		Name:      s.target.Name + "-compute",
	}, &computeStatefulSet)
	if err != nil {
		if apierrors.IsNotFound(err) {
			return nil, nil
		}
		return nil, fmt.Errorf("unable to get state 'computeStatefulSet': %w", err)
	}
	if !ctrlkit.ValidateOwnership(&computeStatefulSet, s.target) {
		return nil, fmt.Errorf("unable to get state 'computeStatefulSet': object not owned by target")
	}

	return &computeStatefulSet, nil
}

// GetCompactorDeployment gets compactorDeployment with name equals to ${target.Name}-compactor.
func (s *RisingWaveControllerManagerState) GetCompactorDeployment(ctx context.Context) (*appsv1.Deployment, error) {
	var compactorDeployment appsv1.Deployment

	err := s.Get(ctx, types.NamespacedName{
		Namespace: s.target.Namespace,
		Name:      s.target.Name + "-compactor",
	}, &compactorDeployment)
	if err != nil {
		if apierrors.IsNotFound(err) {
			return nil, nil
		}
		return nil, fmt.Errorf("unable to get state 'compactorDeployment': %w", err)
	}
	if !ctrlkit.ValidateOwnership(&compactorDeployment, s.target) {
		return nil, fmt.Errorf("unable to get state 'compactorDeployment': object not owned by target")
	}

	return &compactorDeployment, nil
}

// GetFrontendService gets frontendService with name equals to ${target.Name}-frontend.
func (s *RisingWaveControllerManagerState) GetFrontendService(ctx context.Context) (*corev1.Service, error) {
	var frontendService corev1.Service

	err := s.Get(ctx, types.NamespacedName{
		Namespace: s.target.Namespace,
		Name:      s.target.Name + "-frontend",
	}, &frontendService)
	if err != nil {
		if apierrors.IsNotFound(err) {
			return nil, nil
		}
		return nil, fmt.Errorf("unable to get state 'frontendService': %w", err)
	}
	if !ctrlkit.ValidateOwnership(&frontendService, s.target) {
		return nil, fmt.Errorf("unable to get state 'frontendService': object not owned by target")
	}

	return &frontendService, nil
}

// GetComputeService gets computeService with name equals to ${target.Name}-compute.
func (s *RisingWaveControllerManagerState) GetComputeService(ctx context.Context) (*corev1.Service, error) {
	var computeService corev1.Service

	err := s.Get(ctx, types.NamespacedName{
		Namespace: s.target.Namespace,
		Name:      s.target.Name + "-compute",
	}, &computeService)
	if err != nil {
		if apierrors.IsNotFound(err) {
			return nil, nil
		}
		return nil, fmt.Errorf("unable to get state 'computeService': %w", err)
	}
	if !ctrlkit.ValidateOwnership(&computeService, s.target) {
		return nil, fmt.Errorf("unable to get state 'computeService': object not owned by target")
	}

	return &computeService, nil
}

// GetCompactorService gets compactorService with name equals to ${target.Name}-compactor.
func (s *RisingWaveControllerManagerState) GetCompactorService(ctx context.Context) (*corev1.Service, error) {
	var compactorService corev1.Service

	err := s.Get(ctx, types.NamespacedName{
		Namespace: s.target.Namespace,
		Name:      s.target.Name + "-compactor",
	}, &compactorService)
	if err != nil {
		if apierrors.IsNotFound(err) {
			return nil, nil
		}
		return nil, fmt.Errorf("unable to get state 'compactorService': %w", err)
	}
	if !ctrlkit.ValidateOwnership(&compactorService, s.target) {
		return nil, fmt.Errorf("unable to get state 'compactorService': object not owned by target")
	}

	return &compactorService, nil
}

// GetComputeConfigMap gets computeConfigMap with name equals to ${target.Name}-compute.
func (s *RisingWaveControllerManagerState) GetComputeConfigMap(ctx context.Context) (*corev1.ConfigMap, error) {
	var computeConfigMap corev1.ConfigMap

	err := s.Get(ctx, types.NamespacedName{
		Namespace: s.target.Namespace,
		Name:      s.target.Name + "-compute",
	}, &computeConfigMap)
	if err != nil {
		if apierrors.IsNotFound(err) {
			return nil, nil
		}
		return nil, fmt.Errorf("unable to get state 'computeConfigMap': %w", err)
	}
	if !ctrlkit.ValidateOwnership(&computeConfigMap, s.target) {
		return nil, fmt.Errorf("unable to get state 'computeConfigMap': object not owned by target")
	}

	return &computeConfigMap, nil
}

// GetFrontendDeployment gets frontendDeployment with name equals to ${target.Name}-frontend.
func (s *RisingWaveControllerManagerState) GetFrontendDeployment(ctx context.Context) (*appsv1.Deployment, error) {
	var frontendDeployment appsv1.Deployment

	err := s.Get(ctx, types.NamespacedName{
		Namespace: s.target.Namespace,
		Name:      s.target.Name + "-frontend",
	}, &frontendDeployment)
	if err != nil {
		if apierrors.IsNotFound(err) {
			return nil, nil
		}
		return nil, fmt.Errorf("unable to get state 'frontendDeployment': %w", err)
	}
	if !ctrlkit.ValidateOwnership(&frontendDeployment, s.target) {
		return nil, fmt.Errorf("unable to get state 'frontendDeployment': object not owned by target")
	}

	return &frontendDeployment, nil
}

// GetMetaService gets metaService with name equals to ${target.Name}-meta.
func (s *RisingWaveControllerManagerState) GetMetaService(ctx context.Context) (*corev1.Service, error) {
	var metaService corev1.Service

	err := s.Get(ctx, types.NamespacedName{
		Namespace: s.target.Namespace,
		Name:      s.target.Name + "-meta",
	}, &metaService)
	if err != nil {
		if apierrors.IsNotFound(err) {
			return nil, nil
		}
		return nil, fmt.Errorf("unable to get state 'metaService': %w", err)
	}
	if !ctrlkit.ValidateOwnership(&metaService, s.target) {
		return nil, fmt.Errorf("unable to get state 'metaService': object not owned by target")
	}

	return &metaService, nil
}

// GetMetaDeployment gets metaDeployment with name equals to ${target.Name}-meta.
func (s *RisingWaveControllerManagerState) GetMetaDeployment(ctx context.Context) (*appsv1.Deployment, error) {
	var metaDeployment appsv1.Deployment

	err := s.Get(ctx, types.NamespacedName{
		Namespace: s.target.Namespace,
		Name:      s.target.Name + "-meta",
	}, &metaDeployment)
	if err != nil {
		if apierrors.IsNotFound(err) {
			return nil, nil
		}
		return nil, fmt.Errorf("unable to get state 'metaDeployment': %w", err)
	}
	if !ctrlkit.ValidateOwnership(&metaDeployment, s.target) {
		return nil, fmt.Errorf("unable to get state 'metaDeployment': object not owned by target")
	}

	return &metaDeployment, nil
}

// NewRisingWaveControllerManagerState returns a RisingWaveControllerManagerState (target is not copied).
func NewRisingWaveControllerManagerState(reader client.Reader, target *risingwavev1alpha1.RisingWave) RisingWaveControllerManagerState {
	return RisingWaveControllerManagerState{
		Reader: reader,
		target: target,
	}
}

// RisingWaveControllerManagerImpl declares the implementation interface for RisingWaveControllerManager.
type RisingWaveControllerManagerImpl interface {
	ctrlkit.CrontollerManagerActionLifeCycleHook

	// SyncMetaService creates or updates the service for meta nodes.
	SyncMetaService(ctx context.Context, logger logr.Logger, metaService *corev1.Service) (ctrl.Result, error)

	// SyncMetaDeployment creates or updates the deployment for meta nodes.
	SyncMetaDeployment(ctx context.Context, logger logr.Logger, metaDeployment *appsv1.Deployment) (ctrl.Result, error)

	// WaitBeforeMetaServiceIsAvailable waits (aborts the workflow) before the meta service is available.
	WaitBeforeMetaServiceIsAvailable(ctx context.Context, logger logr.Logger, metaService *corev1.Service) (ctrl.Result, error)

	// WaitBeforeMetaDeploymentReady waits (aborts the workflow) before the meta deployment is ready.
	WaitBeforeMetaDeploymentReady(ctx context.Context, logger logr.Logger, metaDeployment *appsv1.Deployment) (ctrl.Result, error)

	// SyncFrontendService creates or updates the service for frontend nodes.
	SyncFrontendService(ctx context.Context, logger logr.Logger, frontendService *corev1.Service) (ctrl.Result, error)

	// SyncFrontendDeployment creates or updates the deployment for frontend nodes.
	SyncFrontendDeployment(ctx context.Context, logger logr.Logger, frontendDeployment *appsv1.Deployment) (ctrl.Result, error)

	// WaitBeforeFrontendDeploymentReady waits (aborts the workflow) before the frontend deployment is ready.
	WaitBeforeFrontendDeploymentReady(ctx context.Context, logger logr.Logger, frontendDeployment *appsv1.Deployment) (ctrl.Result, error)

	// SyncComputeSerivce creates or updates the service for compute nodes.
	SyncComputeSerivce(ctx context.Context, logger logr.Logger, computeService *corev1.Service) (ctrl.Result, error)

	// SyncComputeStatefulSet creates or updates the statefulset for compute nodes.
	SyncComputeStatefulSet(ctx context.Context, logger logr.Logger, computeStatefulSet *appsv1.StatefulSet) (ctrl.Result, error)

	// WaitBeforeComputeStatefulSetReady waits (aborts the workflow) before the compute statefulset is ready.
	WaitBeforeComputeStatefulSetReady(ctx context.Context, logger logr.Logger, computeStatefulSet *appsv1.StatefulSet) (ctrl.Result, error)

	// SyncCompactorService creates or updates the service for compactor nodes.
	SyncCompactorService(ctx context.Context, logger logr.Logger, compactorService *corev1.Service) (ctrl.Result, error)

	// SyncCompactorDeployment creates or updates the deployment for compactor nodes.
	SyncCompactorDeployment(ctx context.Context, logger logr.Logger, compactorDeployment *appsv1.Deployment) (ctrl.Result, error)

	// WaitBeforeCompactorDeploymentReady waits (aborts the workflow) before the compactor deployment is ready.
	WaitBeforeCompactorDeploymentReady(ctx context.Context, logger logr.Logger, compactorDeployment *appsv1.Deployment) (ctrl.Result, error)

	// SyncComputeConfigMap creates or updates the configmap for compute nodes.
	SyncComputeConfigMap(ctx context.Context, logger logr.Logger, computeConfigMap *corev1.ConfigMap) (ctrl.Result, error)

	// CollectRunningStatisticsAndSyncStatus collects running statistics and sync them into the status.
	CollectRunningStatisticsAndSyncStatus(ctx context.Context, logger logr.Logger, frontendService *corev1.Service, metaService *corev1.Service, computeService *corev1.Service, compactorService *corev1.Service, metaDeployment *appsv1.Deployment, frontendDeployment *appsv1.Deployment, computeStatefulSet *appsv1.StatefulSet, compactorDeployment *appsv1.Deployment, computeConfigMap *corev1.ConfigMap) (ctrl.Result, error)
}

// RisingWaveControllerManager encapsulates the states and actions used by RisingWaveController.
type RisingWaveControllerManager struct {
	state  RisingWaveControllerManagerState
	impl   RisingWaveControllerManagerImpl
	logger logr.Logger
}

// WrapAction returns an action from manager.
func (m *RisingWaveControllerManager) WrapAction(description string, f func(context.Context, logr.Logger) (ctrl.Result, error)) ctrlkit.ReconcileAction {
	return ctrlkit.WrapAction(description, func(ctx context.Context) (ctrl.Result, error) {
		logger := m.logger.WithValues("action", description)

		defer m.impl.AfterActionRun(description, ctx, logger)
		m.impl.BeforeActionRun(description, ctx, logger)
		return f(ctx, logger)
	})
}

// SyncMetaService generates the action of "SyncMetaService".
func (m *RisingWaveControllerManager) SyncMetaService() ctrlkit.ReconcileAction {
	return ctrlkit.WrapAction("SyncMetaService", func(ctx context.Context) (ctrl.Result, error) {
		logger := m.logger.WithValues("action", "SyncMetaService")

		// Get states.
		metaService, err := m.state.GetMetaService(ctx)
		if err != nil {
			return ctrlkit.RequeueIfError(err)
		}

		// Invoke action.
		defer m.impl.AfterActionRun("SyncMetaService", ctx, logger)
		m.impl.BeforeActionRun("SyncMetaService", ctx, logger)

		return m.impl.SyncMetaService(ctx, logger, metaService)
	})
}

// SyncMetaDeployment generates the action of "SyncMetaDeployment".
func (m *RisingWaveControllerManager) SyncMetaDeployment() ctrlkit.ReconcileAction {
	return ctrlkit.WrapAction("SyncMetaDeployment", func(ctx context.Context) (ctrl.Result, error) {
		logger := m.logger.WithValues("action", "SyncMetaDeployment")

		// Get states.
		metaDeployment, err := m.state.GetMetaDeployment(ctx)
		if err != nil {
			return ctrlkit.RequeueIfError(err)
		}

		// Invoke action.
		defer m.impl.AfterActionRun("SyncMetaDeployment", ctx, logger)
		m.impl.BeforeActionRun("SyncMetaDeployment", ctx, logger)

		return m.impl.SyncMetaDeployment(ctx, logger, metaDeployment)
	})
}

// WaitBeforeMetaServiceIsAvailable generates the action of "WaitBeforeMetaServiceIsAvailable".
func (m *RisingWaveControllerManager) WaitBeforeMetaServiceIsAvailable() ctrlkit.ReconcileAction {
	return ctrlkit.WrapAction("WaitBeforeMetaServiceIsAvailable", func(ctx context.Context) (ctrl.Result, error) {
		logger := m.logger.WithValues("action", "WaitBeforeMetaServiceIsAvailable")

		// Get states.
		metaService, err := m.state.GetMetaService(ctx)
		if err != nil {
			return ctrlkit.RequeueIfError(err)
		}

		// Invoke action.
		defer m.impl.AfterActionRun("WaitBeforeMetaServiceIsAvailable", ctx, logger)
		m.impl.BeforeActionRun("WaitBeforeMetaServiceIsAvailable", ctx, logger)

		return m.impl.WaitBeforeMetaServiceIsAvailable(ctx, logger, metaService)
	})
}

// WaitBeforeMetaDeploymentReady generates the action of "WaitBeforeMetaDeploymentReady".
func (m *RisingWaveControllerManager) WaitBeforeMetaDeploymentReady() ctrlkit.ReconcileAction {
	return ctrlkit.WrapAction("WaitBeforeMetaDeploymentReady", func(ctx context.Context) (ctrl.Result, error) {
		logger := m.logger.WithValues("action", "WaitBeforeMetaDeploymentReady")

		// Get states.
		metaDeployment, err := m.state.GetMetaDeployment(ctx)
		if err != nil {
			return ctrlkit.RequeueIfError(err)
		}

		// Invoke action.
		defer m.impl.AfterActionRun("WaitBeforeMetaDeploymentReady", ctx, logger)
		m.impl.BeforeActionRun("WaitBeforeMetaDeploymentReady", ctx, logger)

		return m.impl.WaitBeforeMetaDeploymentReady(ctx, logger, metaDeployment)
	})
}

// SyncFrontendService generates the action of "SyncFrontendService".
func (m *RisingWaveControllerManager) SyncFrontendService() ctrlkit.ReconcileAction {
	return ctrlkit.WrapAction("SyncFrontendService", func(ctx context.Context) (ctrl.Result, error) {
		logger := m.logger.WithValues("action", "SyncFrontendService")

		// Get states.
		frontendService, err := m.state.GetFrontendService(ctx)
		if err != nil {
			return ctrlkit.RequeueIfError(err)
		}

		// Invoke action.
		defer m.impl.AfterActionRun("SyncFrontendService", ctx, logger)
		m.impl.BeforeActionRun("SyncFrontendService", ctx, logger)

		return m.impl.SyncFrontendService(ctx, logger, frontendService)
	})
}

// SyncFrontendDeployment generates the action of "SyncFrontendDeployment".
func (m *RisingWaveControllerManager) SyncFrontendDeployment() ctrlkit.ReconcileAction {
	return ctrlkit.WrapAction("SyncFrontendDeployment", func(ctx context.Context) (ctrl.Result, error) {
		logger := m.logger.WithValues("action", "SyncFrontendDeployment")

		// Get states.
		frontendDeployment, err := m.state.GetFrontendDeployment(ctx)
		if err != nil {
			return ctrlkit.RequeueIfError(err)
		}

		// Invoke action.
		defer m.impl.AfterActionRun("SyncFrontendDeployment", ctx, logger)
		m.impl.BeforeActionRun("SyncFrontendDeployment", ctx, logger)

		return m.impl.SyncFrontendDeployment(ctx, logger, frontendDeployment)
	})
}

// WaitBeforeFrontendDeploymentReady generates the action of "WaitBeforeFrontendDeploymentReady".
func (m *RisingWaveControllerManager) WaitBeforeFrontendDeploymentReady() ctrlkit.ReconcileAction {
	return ctrlkit.WrapAction("WaitBeforeFrontendDeploymentReady", func(ctx context.Context) (ctrl.Result, error) {
		logger := m.logger.WithValues("action", "WaitBeforeFrontendDeploymentReady")

		// Get states.
		frontendDeployment, err := m.state.GetFrontendDeployment(ctx)
		if err != nil {
			return ctrlkit.RequeueIfError(err)
		}

		// Invoke action.
		defer m.impl.AfterActionRun("WaitBeforeFrontendDeploymentReady", ctx, logger)
		m.impl.BeforeActionRun("WaitBeforeFrontendDeploymentReady", ctx, logger)

		return m.impl.WaitBeforeFrontendDeploymentReady(ctx, logger, frontendDeployment)
	})
}

// SyncComputeSerivce generates the action of "SyncComputeSerivce".
func (m *RisingWaveControllerManager) SyncComputeSerivce() ctrlkit.ReconcileAction {
	return ctrlkit.WrapAction("SyncComputeSerivce", func(ctx context.Context) (ctrl.Result, error) {
		logger := m.logger.WithValues("action", "SyncComputeSerivce")

		// Get states.
		computeService, err := m.state.GetComputeService(ctx)
		if err != nil {
			return ctrlkit.RequeueIfError(err)
		}

		// Invoke action.
		defer m.impl.AfterActionRun("SyncComputeSerivce", ctx, logger)
		m.impl.BeforeActionRun("SyncComputeSerivce", ctx, logger)

		return m.impl.SyncComputeSerivce(ctx, logger, computeService)
	})
}

// SyncComputeStatefulSet generates the action of "SyncComputeStatefulSet".
func (m *RisingWaveControllerManager) SyncComputeStatefulSet() ctrlkit.ReconcileAction {
	return ctrlkit.WrapAction("SyncComputeStatefulSet", func(ctx context.Context) (ctrl.Result, error) {
		logger := m.logger.WithValues("action", "SyncComputeStatefulSet")

		// Get states.
		computeStatefulSet, err := m.state.GetComputeStatefulSet(ctx)
		if err != nil {
			return ctrlkit.RequeueIfError(err)
		}

		// Invoke action.
		defer m.impl.AfterActionRun("SyncComputeStatefulSet", ctx, logger)
		m.impl.BeforeActionRun("SyncComputeStatefulSet", ctx, logger)

		return m.impl.SyncComputeStatefulSet(ctx, logger, computeStatefulSet)
	})
}

// WaitBeforeComputeStatefulSetReady generates the action of "WaitBeforeComputeStatefulSetReady".
func (m *RisingWaveControllerManager) WaitBeforeComputeStatefulSetReady() ctrlkit.ReconcileAction {
	return ctrlkit.WrapAction("WaitBeforeComputeStatefulSetReady", func(ctx context.Context) (ctrl.Result, error) {
		logger := m.logger.WithValues("action", "WaitBeforeComputeStatefulSetReady")

		// Get states.
		computeStatefulSet, err := m.state.GetComputeStatefulSet(ctx)
		if err != nil {
			return ctrlkit.RequeueIfError(err)
		}

		// Invoke action.
		defer m.impl.AfterActionRun("WaitBeforeComputeStatefulSetReady", ctx, logger)
		m.impl.BeforeActionRun("WaitBeforeComputeStatefulSetReady", ctx, logger)

		return m.impl.WaitBeforeComputeStatefulSetReady(ctx, logger, computeStatefulSet)
	})
}

// SyncCompactorService generates the action of "SyncCompactorService".
func (m *RisingWaveControllerManager) SyncCompactorService() ctrlkit.ReconcileAction {
	return ctrlkit.WrapAction("SyncCompactorService", func(ctx context.Context) (ctrl.Result, error) {
		logger := m.logger.WithValues("action", "SyncCompactorService")

		// Get states.
		compactorService, err := m.state.GetCompactorService(ctx)
		if err != nil {
			return ctrlkit.RequeueIfError(err)
		}

		// Invoke action.
		defer m.impl.AfterActionRun("SyncCompactorService", ctx, logger)
		m.impl.BeforeActionRun("SyncCompactorService", ctx, logger)

		return m.impl.SyncCompactorService(ctx, logger, compactorService)
	})
}

// SyncCompactorDeployment generates the action of "SyncCompactorDeployment".
func (m *RisingWaveControllerManager) SyncCompactorDeployment() ctrlkit.ReconcileAction {
	return ctrlkit.WrapAction("SyncCompactorDeployment", func(ctx context.Context) (ctrl.Result, error) {
		logger := m.logger.WithValues("action", "SyncCompactorDeployment")

		// Get states.
		compactorDeployment, err := m.state.GetCompactorDeployment(ctx)
		if err != nil {
			return ctrlkit.RequeueIfError(err)
		}

		// Invoke action.
		defer m.impl.AfterActionRun("SyncCompactorDeployment", ctx, logger)
		m.impl.BeforeActionRun("SyncCompactorDeployment", ctx, logger)

		return m.impl.SyncCompactorDeployment(ctx, logger, compactorDeployment)
	})
}

// WaitBeforeCompactorDeploymentReady generates the action of "WaitBeforeCompactorDeploymentReady".
func (m *RisingWaveControllerManager) WaitBeforeCompactorDeploymentReady() ctrlkit.ReconcileAction {
	return ctrlkit.WrapAction("WaitBeforeCompactorDeploymentReady", func(ctx context.Context) (ctrl.Result, error) {
		logger := m.logger.WithValues("action", "WaitBeforeCompactorDeploymentReady")

		// Get states.
		compactorDeployment, err := m.state.GetCompactorDeployment(ctx)
		if err != nil {
			return ctrlkit.RequeueIfError(err)
		}

		// Invoke action.
		defer m.impl.AfterActionRun("WaitBeforeCompactorDeploymentReady", ctx, logger)
		m.impl.BeforeActionRun("WaitBeforeCompactorDeploymentReady", ctx, logger)

		return m.impl.WaitBeforeCompactorDeploymentReady(ctx, logger, compactorDeployment)
	})
}

// SyncComputeConfigMap generates the action of "SyncComputeConfigMap".
func (m *RisingWaveControllerManager) SyncComputeConfigMap() ctrlkit.ReconcileAction {
	return ctrlkit.WrapAction("SyncComputeConfigMap", func(ctx context.Context) (ctrl.Result, error) {
		logger := m.logger.WithValues("action", "SyncComputeConfigMap")

		// Get states.
		computeConfigMap, err := m.state.GetComputeConfigMap(ctx)
		if err != nil {
			return ctrlkit.RequeueIfError(err)
		}

		// Invoke action.
		defer m.impl.AfterActionRun("SyncComputeConfigMap", ctx, logger)
		m.impl.BeforeActionRun("SyncComputeConfigMap", ctx, logger)

		return m.impl.SyncComputeConfigMap(ctx, logger, computeConfigMap)
	})
}

// CollectRunningStatisticsAndSyncStatus generates the action of "CollectRunningStatisticsAndSyncStatus".
func (m *RisingWaveControllerManager) CollectRunningStatisticsAndSyncStatus() ctrlkit.ReconcileAction {
	return ctrlkit.WrapAction("CollectRunningStatisticsAndSyncStatus", func(ctx context.Context) (ctrl.Result, error) {
		logger := m.logger.WithValues("action", "CollectRunningStatisticsAndSyncStatus")

		// Get states.
		frontendService, err := m.state.GetFrontendService(ctx)
		if err != nil {
			return ctrlkit.RequeueIfError(err)
		}

		metaService, err := m.state.GetMetaService(ctx)
		if err != nil {
			return ctrlkit.RequeueIfError(err)
		}

		computeService, err := m.state.GetComputeService(ctx)
		if err != nil {
			return ctrlkit.RequeueIfError(err)
		}

		compactorService, err := m.state.GetCompactorService(ctx)
		if err != nil {
			return ctrlkit.RequeueIfError(err)
		}

		metaDeployment, err := m.state.GetMetaDeployment(ctx)
		if err != nil {
			return ctrlkit.RequeueIfError(err)
		}

		frontendDeployment, err := m.state.GetFrontendDeployment(ctx)
		if err != nil {
			return ctrlkit.RequeueIfError(err)
		}

		computeStatefulSet, err := m.state.GetComputeStatefulSet(ctx)
		if err != nil {
			return ctrlkit.RequeueIfError(err)
		}

		compactorDeployment, err := m.state.GetCompactorDeployment(ctx)
		if err != nil {
			return ctrlkit.RequeueIfError(err)
		}

		computeConfigMap, err := m.state.GetComputeConfigMap(ctx)
		if err != nil {
			return ctrlkit.RequeueIfError(err)
		}

		// Invoke action.
		defer m.impl.AfterActionRun("CollectRunningStatisticsAndSyncStatus", ctx, logger)
		m.impl.BeforeActionRun("CollectRunningStatisticsAndSyncStatus", ctx, logger)

		return m.impl.CollectRunningStatisticsAndSyncStatus(ctx, logger, frontendService, metaService, computeService, compactorService, metaDeployment, frontendDeployment, computeStatefulSet, compactorDeployment, computeConfigMap)
	})
}

// NewRisingWaveControllerManager returns a new RisingWaveControllerManager with given state and implementation.
func NewRisingWaveControllerManager(state RisingWaveControllerManagerState, impl RisingWaveControllerManagerImpl, logger logr.Logger) RisingWaveControllerManager {
	return RisingWaveControllerManager{
		state:  state,
		impl:   impl,
		logger: logger,
	}
}
