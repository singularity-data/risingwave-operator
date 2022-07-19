//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by ctrlkit. DO NOT EDIT.

package manager

import (
	"context"
	"fmt"

	"github.com/go-logr/logr"
	monitoringv1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
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

// GetCompactorDeployments lists compactorDeployments with the following selectors:
//   + labels/risingwave/component=compactor
//   + labels/risingwave/name=${target.Name}
//   + owned
func (s *RisingWaveControllerManagerState) GetCompactorDeployments(ctx context.Context) ([]appsv1.Deployment, error) {
	var compactorDeploymentsList appsv1.DeploymentList

	matchingLabels := map[string]string{
		"risingwave/component": "compactor",
		"risingwave/name":      s.target.Name,
	}

	err := s.List(ctx, &compactorDeploymentsList, client.InNamespace(s.target.Namespace),
		client.MatchingLabels(matchingLabels))
	if err != nil {
		return nil, fmt.Errorf("unable to get state 'compactorDeployments': %w", err)
	}

	var validated []appsv1.Deployment
	for _, obj := range compactorDeploymentsList.Items {
		if ctrlkit.ValidateOwnership(&obj, s.target) {
			validated = append(validated, obj)
		}
	}

	return validated, nil
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

// GetComputeStatefulSets lists computeStatefulSets with the following selectors:
//   + labels/risingwave/component=compute
//   + labels/risingwave/name=${target.Name}
//   + owned
func (s *RisingWaveControllerManagerState) GetComputeStatefulSets(ctx context.Context) ([]appsv1.StatefulSet, error) {
	var computeStatefulSetsList appsv1.StatefulSetList

	matchingLabels := map[string]string{
		"risingwave/component": "compute",
		"risingwave/name":      s.target.Name,
	}

	err := s.List(ctx, &computeStatefulSetsList, client.InNamespace(s.target.Namespace),
		client.MatchingLabels(matchingLabels))
	if err != nil {
		return nil, fmt.Errorf("unable to get state 'computeStatefulSets': %w", err)
	}

	var validated []appsv1.StatefulSet
	for _, obj := range computeStatefulSetsList.Items {
		if ctrlkit.ValidateOwnership(&obj, s.target) {
			validated = append(validated, obj)
		}
	}

	return validated, nil
}

// GetConfigConfigMap gets configConfigMap with name equals to ${target.Name}-config.
func (s *RisingWaveControllerManagerState) GetConfigConfigMap(ctx context.Context) (*corev1.ConfigMap, error) {
	var configConfigMap corev1.ConfigMap

	err := s.Get(ctx, types.NamespacedName{
		Namespace: s.target.Namespace,
		Name:      s.target.Name + "-config",
	}, &configConfigMap)
	if err != nil {
		if apierrors.IsNotFound(err) {
			return nil, nil
		}
		return nil, fmt.Errorf("unable to get state 'configConfigMap': %w", err)
	}
	if !ctrlkit.ValidateOwnership(&configConfigMap, s.target) {
		return nil, fmt.Errorf("unable to get state 'configConfigMap': object not owned by target")
	}

	return &configConfigMap, nil
}

// GetFrontendDeployments lists frontendDeployments with the following selectors:
//   + labels/risingwave/component=frontend
//   + labels/risingwave/name=${target.Name}
//   + owned
func (s *RisingWaveControllerManagerState) GetFrontendDeployments(ctx context.Context) ([]appsv1.Deployment, error) {
	var frontendDeploymentsList appsv1.DeploymentList

	matchingLabels := map[string]string{
		"risingwave/component": "frontend",
		"risingwave/name":      s.target.Name,
	}

	err := s.List(ctx, &frontendDeploymentsList, client.InNamespace(s.target.Namespace),
		client.MatchingLabels(matchingLabels))
	if err != nil {
		return nil, fmt.Errorf("unable to get state 'frontendDeployments': %w", err)
	}

	var validated []appsv1.Deployment
	for _, obj := range frontendDeploymentsList.Items {
		if ctrlkit.ValidateOwnership(&obj, s.target) {
			validated = append(validated, obj)
		}
	}

	return validated, nil
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

// GetMetaDeployments lists metaDeployments with the following selectors:
//   + labels/risingwave/component=meta
//   + labels/risingwave/name=${target.Name}
//   + owned
func (s *RisingWaveControllerManagerState) GetMetaDeployments(ctx context.Context) ([]appsv1.Deployment, error) {
	var metaDeploymentsList appsv1.DeploymentList

	matchingLabels := map[string]string{
		"risingwave/component": "meta",
		"risingwave/name":      s.target.Name,
	}

	err := s.List(ctx, &metaDeploymentsList, client.InNamespace(s.target.Namespace),
		client.MatchingLabels(matchingLabels))
	if err != nil {
		return nil, fmt.Errorf("unable to get state 'metaDeployments': %w", err)
	}

	var validated []appsv1.Deployment
	for _, obj := range metaDeploymentsList.Items {
		if ctrlkit.ValidateOwnership(&obj, s.target) {
			validated = append(validated, obj)
		}
	}

	return validated, nil
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

// GetServiceMonitor gets serviceMonitor with name equals to risingwave-${target.Name}.
func (s *RisingWaveControllerManagerState) GetServiceMonitor(ctx context.Context) (*monitoringv1.ServiceMonitor, error) {
	var serviceMonitor monitoringv1.ServiceMonitor

	err := s.Get(ctx, types.NamespacedName{
		Namespace: s.target.Namespace,
		Name:      "risingwave-" + s.target.Name,
	}, &serviceMonitor)
	if err != nil {
		if apierrors.IsNotFound(err) {
			return nil, nil
		}
		return nil, fmt.Errorf("unable to get state 'serviceMonitor': %w", err)
	}
	if !ctrlkit.ValidateOwnership(&serviceMonitor, s.target) {
		return nil, fmt.Errorf("unable to get state 'serviceMonitor': object not owned by target")
	}

	return &serviceMonitor, nil
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
	// SyncMetaService creates or updates the service for meta nodes.
	SyncMetaService(ctx context.Context, logger logr.Logger, metaService *corev1.Service) (ctrl.Result, error)

	// SyncMetaDeployments creates or updates the deployments for meta nodes.
	SyncMetaDeployments(ctx context.Context, logger logr.Logger, metaDeployments []appsv1.Deployment) (ctrl.Result, error)

	// WaitBeforeMetaServiceIsAvailable waits (aborts the workflow) before the meta service is available.
	WaitBeforeMetaServiceIsAvailable(ctx context.Context, logger logr.Logger, metaService *corev1.Service) (ctrl.Result, error)

	// WaitBeforeMetaDeploymentsReady waits (aborts the workflow) before the meta deployments are ready.
	WaitBeforeMetaDeploymentsReady(ctx context.Context, logger logr.Logger, metaDeployments []appsv1.Deployment) (ctrl.Result, error)

	// SyncFrontendService creates or updates the service for frontend nodes.
	SyncFrontendService(ctx context.Context, logger logr.Logger, frontendService *corev1.Service) (ctrl.Result, error)

	// SyncFrontendDeployments creates or updates the deployments for frontend nodes.
	SyncFrontendDeployments(ctx context.Context, logger logr.Logger, frontendDeployments []appsv1.Deployment) (ctrl.Result, error)

	// WaitBeforeFrontendDeploymentsReady waits (aborts the workflow) before the frontend deployments are ready.
	WaitBeforeFrontendDeploymentsReady(ctx context.Context, logger logr.Logger, frontendDeployments []appsv1.Deployment) (ctrl.Result, error)

	// SyncComputeService creates or updates the service for compute nodes.
	SyncComputeService(ctx context.Context, logger logr.Logger, computeService *corev1.Service) (ctrl.Result, error)

	// SyncComputeStatefulSets creates or updates the statefulsets for compute nodes.
	SyncComputeStatefulSets(ctx context.Context, logger logr.Logger, computeStatefulSets []appsv1.StatefulSet) (ctrl.Result, error)

	// WaitBeforeComputeStatefulSetsReady waits (aborts the workflow) before the compute statefulsets are ready.
	WaitBeforeComputeStatefulSetsReady(ctx context.Context, logger logr.Logger, computeStatefulSets []appsv1.StatefulSet) (ctrl.Result, error)

	// SyncCompactorService creates or updates the service for compactor nodes.
	SyncCompactorService(ctx context.Context, logger logr.Logger, compactorService *corev1.Service) (ctrl.Result, error)

	// SyncCompactorDeployments creates or updates the deployments for compactor nodes.
	SyncCompactorDeployments(ctx context.Context, logger logr.Logger, compactorDeployments []appsv1.Deployment) (ctrl.Result, error)

	// WaitBeforeCompactorDeploymentsReady waits (aborts the workflow) before the compactor deployments are ready.
	WaitBeforeCompactorDeploymentsReady(ctx context.Context, logger logr.Logger, compactorDeployments []appsv1.Deployment) (ctrl.Result, error)

	// SyncConfigConfigMap creates or updates the configmap for RisingWave configs.
	SyncConfigConfigMap(ctx context.Context, logger logr.Logger, configConfigMap *corev1.ConfigMap) (ctrl.Result, error)

	// SyncServiceMonitor creates or updates the service monitor for RisingWave.
	SyncServiceMonitor(ctx context.Context, logger logr.Logger, serviceMonitor *monitoringv1.ServiceMonitor) (ctrl.Result, error)

	// CollectRunningStatisticsAndSyncStatus collects running statistics and sync them into the status.
	CollectRunningStatisticsAndSyncStatus(ctx context.Context, logger logr.Logger, frontendService *corev1.Service, metaService *corev1.Service, computeService *corev1.Service, compactorService *corev1.Service, metaDeployments []appsv1.Deployment, frontendDeployments []appsv1.Deployment, computeStatefulSets []appsv1.StatefulSet, compactorDeployments []appsv1.Deployment, configConfigMap *corev1.ConfigMap, serviceMonitor *monitoringv1.ServiceMonitor) (ctrl.Result, error)
}

// Pre-defined actions in RisingWaveControllerManager.
const (
	RisingWaveAction_SyncMetaService                       = "SyncMetaService"
	RisingWaveAction_SyncMetaDeployments                   = "SyncMetaDeployments"
	RisingWaveAction_WaitBeforeMetaServiceIsAvailable      = "WaitBeforeMetaServiceIsAvailable"
	RisingWaveAction_WaitBeforeMetaDeploymentsReady        = "WaitBeforeMetaDeploymentsReady"
	RisingWaveAction_SyncFrontendService                   = "SyncFrontendService"
	RisingWaveAction_SyncFrontendDeployments               = "SyncFrontendDeployments"
	RisingWaveAction_WaitBeforeFrontendDeploymentsReady    = "WaitBeforeFrontendDeploymentsReady"
	RisingWaveAction_SyncComputeService                    = "SyncComputeService"
	RisingWaveAction_SyncComputeStatefulSets               = "SyncComputeStatefulSets"
	RisingWaveAction_WaitBeforeComputeStatefulSetsReady    = "WaitBeforeComputeStatefulSetsReady"
	RisingWaveAction_SyncCompactorService                  = "SyncCompactorService"
	RisingWaveAction_SyncCompactorDeployments              = "SyncCompactorDeployments"
	RisingWaveAction_WaitBeforeCompactorDeploymentsReady   = "WaitBeforeCompactorDeploymentsReady"
	RisingWaveAction_SyncConfigConfigMap                   = "SyncConfigConfigMap"
	RisingWaveAction_SyncServiceMonitor                    = "SyncServiceMonitor"
	RisingWaveAction_CollectRunningStatisticsAndSyncStatus = "CollectRunningStatisticsAndSyncStatus"
)

// RisingWaveControllerManager encapsulates the states and actions used by RisingWaveController.
type RisingWaveControllerManager struct {
	hook   ctrlkit.ActionHook
	state  RisingWaveControllerManagerState
	impl   RisingWaveControllerManagerImpl
	logger logr.Logger
}

// NewAction returns a new action controlled by the manager.
func (m *RisingWaveControllerManager) NewAction(description string, f func(context.Context, logr.Logger) (ctrl.Result, error)) ctrlkit.Action {
	return ctrlkit.NewAction(description, func(ctx context.Context) (result ctrl.Result, err error) {
		logger := m.logger.WithValues("action", description)

		if m.hook != nil {
			defer func() { m.hook.PostRun(ctx, logger, description, result, err) }()
			m.hook.PreRun(ctx, logger, description, nil)
		}

		return f(ctx, logger)
	})
}

// SyncMetaService generates the action of "SyncMetaService".
func (m *RisingWaveControllerManager) SyncMetaService() ctrlkit.Action {
	return ctrlkit.NewAction(RisingWaveAction_SyncMetaService, func(ctx context.Context) (result ctrl.Result, err error) {
		logger := m.logger.WithValues("action", RisingWaveAction_SyncMetaService)

		// Get states.
		metaService, err := m.state.GetMetaService(ctx)
		if err != nil {
			return ctrlkit.RequeueIfError(err)
		}

		// Invoke action.
		if m.hook != nil {
			defer func() { m.hook.PostRun(ctx, logger, RisingWaveAction_SyncMetaService, result, err) }()
			m.hook.PreRun(ctx, logger, RisingWaveAction_SyncMetaService, map[string]runtime.Object{
				"metaService": metaService,
			})
		}

		return m.impl.SyncMetaService(ctx, logger, metaService)
	})
}

// SyncMetaDeployments generates the action of "SyncMetaDeployments".
func (m *RisingWaveControllerManager) SyncMetaDeployments() ctrlkit.Action {
	return ctrlkit.NewAction(RisingWaveAction_SyncMetaDeployments, func(ctx context.Context) (result ctrl.Result, err error) {
		logger := m.logger.WithValues("action", RisingWaveAction_SyncMetaDeployments)

		// Get states.
		metaDeployments, err := m.state.GetMetaDeployments(ctx)
		if err != nil {
			return ctrlkit.RequeueIfError(err)
		}

		// Invoke action.
		if m.hook != nil {
			defer func() { m.hook.PostRun(ctx, logger, RisingWaveAction_SyncMetaDeployments, result, err) }()
			m.hook.PreRun(ctx, logger, RisingWaveAction_SyncMetaDeployments, map[string]runtime.Object{
				"metaDeployments": &appsv1.DeploymentList{Items: metaDeployments},
			})
		}

		return m.impl.SyncMetaDeployments(ctx, logger, metaDeployments)
	})
}

// WaitBeforeMetaServiceIsAvailable generates the action of "WaitBeforeMetaServiceIsAvailable".
func (m *RisingWaveControllerManager) WaitBeforeMetaServiceIsAvailable() ctrlkit.Action {
	return ctrlkit.NewAction(RisingWaveAction_WaitBeforeMetaServiceIsAvailable, func(ctx context.Context) (result ctrl.Result, err error) {
		logger := m.logger.WithValues("action", RisingWaveAction_WaitBeforeMetaServiceIsAvailable)

		// Get states.
		metaService, err := m.state.GetMetaService(ctx)
		if err != nil {
			return ctrlkit.RequeueIfError(err)
		}

		// Invoke action.
		if m.hook != nil {
			defer func() { m.hook.PostRun(ctx, logger, RisingWaveAction_WaitBeforeMetaServiceIsAvailable, result, err) }()
			m.hook.PreRun(ctx, logger, RisingWaveAction_WaitBeforeMetaServiceIsAvailable, map[string]runtime.Object{
				"metaService": metaService,
			})
		}

		return m.impl.WaitBeforeMetaServiceIsAvailable(ctx, logger, metaService)
	})
}

// WaitBeforeMetaDeploymentsReady generates the action of "WaitBeforeMetaDeploymentsReady".
func (m *RisingWaveControllerManager) WaitBeforeMetaDeploymentsReady() ctrlkit.Action {
	return ctrlkit.NewAction(RisingWaveAction_WaitBeforeMetaDeploymentsReady, func(ctx context.Context) (result ctrl.Result, err error) {
		logger := m.logger.WithValues("action", RisingWaveAction_WaitBeforeMetaDeploymentsReady)

		// Get states.
		metaDeployments, err := m.state.GetMetaDeployments(ctx)
		if err != nil {
			return ctrlkit.RequeueIfError(err)
		}

		// Invoke action.
		if m.hook != nil {
			defer func() { m.hook.PostRun(ctx, logger, RisingWaveAction_WaitBeforeMetaDeploymentsReady, result, err) }()
			m.hook.PreRun(ctx, logger, RisingWaveAction_WaitBeforeMetaDeploymentsReady, map[string]runtime.Object{
				"metaDeployments": &appsv1.DeploymentList{Items: metaDeployments},
			})
		}

		return m.impl.WaitBeforeMetaDeploymentsReady(ctx, logger, metaDeployments)
	})
}

// SyncFrontendService generates the action of "SyncFrontendService".
func (m *RisingWaveControllerManager) SyncFrontendService() ctrlkit.Action {
	return ctrlkit.NewAction(RisingWaveAction_SyncFrontendService, func(ctx context.Context) (result ctrl.Result, err error) {
		logger := m.logger.WithValues("action", RisingWaveAction_SyncFrontendService)

		// Get states.
		frontendService, err := m.state.GetFrontendService(ctx)
		if err != nil {
			return ctrlkit.RequeueIfError(err)
		}

		// Invoke action.
		if m.hook != nil {
			defer func() { m.hook.PostRun(ctx, logger, RisingWaveAction_SyncFrontendService, result, err) }()
			m.hook.PreRun(ctx, logger, RisingWaveAction_SyncFrontendService, map[string]runtime.Object{
				"frontendService": frontendService,
			})
		}

		return m.impl.SyncFrontendService(ctx, logger, frontendService)
	})
}

// SyncFrontendDeployments generates the action of "SyncFrontendDeployments".
func (m *RisingWaveControllerManager) SyncFrontendDeployments() ctrlkit.Action {
	return ctrlkit.NewAction(RisingWaveAction_SyncFrontendDeployments, func(ctx context.Context) (result ctrl.Result, err error) {
		logger := m.logger.WithValues("action", RisingWaveAction_SyncFrontendDeployments)

		// Get states.
		frontendDeployments, err := m.state.GetFrontendDeployments(ctx)
		if err != nil {
			return ctrlkit.RequeueIfError(err)
		}

		// Invoke action.
		if m.hook != nil {
			defer func() { m.hook.PostRun(ctx, logger, RisingWaveAction_SyncFrontendDeployments, result, err) }()
			m.hook.PreRun(ctx, logger, RisingWaveAction_SyncFrontendDeployments, map[string]runtime.Object{
				"frontendDeployments": &appsv1.DeploymentList{Items: frontendDeployments},
			})
		}

		return m.impl.SyncFrontendDeployments(ctx, logger, frontendDeployments)
	})
}

// WaitBeforeFrontendDeploymentsReady generates the action of "WaitBeforeFrontendDeploymentsReady".
func (m *RisingWaveControllerManager) WaitBeforeFrontendDeploymentsReady() ctrlkit.Action {
	return ctrlkit.NewAction(RisingWaveAction_WaitBeforeFrontendDeploymentsReady, func(ctx context.Context) (result ctrl.Result, err error) {
		logger := m.logger.WithValues("action", RisingWaveAction_WaitBeforeFrontendDeploymentsReady)

		// Get states.
		frontendDeployments, err := m.state.GetFrontendDeployments(ctx)
		if err != nil {
			return ctrlkit.RequeueIfError(err)
		}

		// Invoke action.
		if m.hook != nil {
			defer func() { m.hook.PostRun(ctx, logger, RisingWaveAction_WaitBeforeFrontendDeploymentsReady, result, err) }()
			m.hook.PreRun(ctx, logger, RisingWaveAction_WaitBeforeFrontendDeploymentsReady, map[string]runtime.Object{
				"frontendDeployments": &appsv1.DeploymentList{Items: frontendDeployments},
			})
		}

		return m.impl.WaitBeforeFrontendDeploymentsReady(ctx, logger, frontendDeployments)
	})
}

// SyncComputeService generates the action of "SyncComputeService".
func (m *RisingWaveControllerManager) SyncComputeService() ctrlkit.Action {
	return ctrlkit.NewAction(RisingWaveAction_SyncComputeService, func(ctx context.Context) (result ctrl.Result, err error) {
		logger := m.logger.WithValues("action", RisingWaveAction_SyncComputeService)

		// Get states.
		computeService, err := m.state.GetComputeService(ctx)
		if err != nil {
			return ctrlkit.RequeueIfError(err)
		}

		// Invoke action.
		if m.hook != nil {
			defer func() { m.hook.PostRun(ctx, logger, RisingWaveAction_SyncComputeService, result, err) }()
			m.hook.PreRun(ctx, logger, RisingWaveAction_SyncComputeService, map[string]runtime.Object{
				"computeService": computeService,
			})
		}

		return m.impl.SyncComputeService(ctx, logger, computeService)
	})
}

// SyncComputeStatefulSets generates the action of "SyncComputeStatefulSets".
func (m *RisingWaveControllerManager) SyncComputeStatefulSets() ctrlkit.Action {
	return ctrlkit.NewAction(RisingWaveAction_SyncComputeStatefulSets, func(ctx context.Context) (result ctrl.Result, err error) {
		logger := m.logger.WithValues("action", RisingWaveAction_SyncComputeStatefulSets)

		// Get states.
		computeStatefulSets, err := m.state.GetComputeStatefulSets(ctx)
		if err != nil {
			return ctrlkit.RequeueIfError(err)
		}

		// Invoke action.
		if m.hook != nil {
			defer func() { m.hook.PostRun(ctx, logger, RisingWaveAction_SyncComputeStatefulSets, result, err) }()
			m.hook.PreRun(ctx, logger, RisingWaveAction_SyncComputeStatefulSets, map[string]runtime.Object{
				"computeStatefulSets": &appsv1.StatefulSetList{Items: computeStatefulSets},
			})
		}

		return m.impl.SyncComputeStatefulSets(ctx, logger, computeStatefulSets)
	})
}

// WaitBeforeComputeStatefulSetsReady generates the action of "WaitBeforeComputeStatefulSetsReady".
func (m *RisingWaveControllerManager) WaitBeforeComputeStatefulSetsReady() ctrlkit.Action {
	return ctrlkit.NewAction(RisingWaveAction_WaitBeforeComputeStatefulSetsReady, func(ctx context.Context) (result ctrl.Result, err error) {
		logger := m.logger.WithValues("action", RisingWaveAction_WaitBeforeComputeStatefulSetsReady)

		// Get states.
		computeStatefulSets, err := m.state.GetComputeStatefulSets(ctx)
		if err != nil {
			return ctrlkit.RequeueIfError(err)
		}

		// Invoke action.
		if m.hook != nil {
			defer func() { m.hook.PostRun(ctx, logger, RisingWaveAction_WaitBeforeComputeStatefulSetsReady, result, err) }()
			m.hook.PreRun(ctx, logger, RisingWaveAction_WaitBeforeComputeStatefulSetsReady, map[string]runtime.Object{
				"computeStatefulSets": &appsv1.StatefulSetList{Items: computeStatefulSets},
			})
		}

		return m.impl.WaitBeforeComputeStatefulSetsReady(ctx, logger, computeStatefulSets)
	})
}

// SyncCompactorService generates the action of "SyncCompactorService".
func (m *RisingWaveControllerManager) SyncCompactorService() ctrlkit.Action {
	return ctrlkit.NewAction(RisingWaveAction_SyncCompactorService, func(ctx context.Context) (result ctrl.Result, err error) {
		logger := m.logger.WithValues("action", RisingWaveAction_SyncCompactorService)

		// Get states.
		compactorService, err := m.state.GetCompactorService(ctx)
		if err != nil {
			return ctrlkit.RequeueIfError(err)
		}

		// Invoke action.
		if m.hook != nil {
			defer func() { m.hook.PostRun(ctx, logger, RisingWaveAction_SyncCompactorService, result, err) }()
			m.hook.PreRun(ctx, logger, RisingWaveAction_SyncCompactorService, map[string]runtime.Object{
				"compactorService": compactorService,
			})
		}

		return m.impl.SyncCompactorService(ctx, logger, compactorService)
	})
}

// SyncCompactorDeployments generates the action of "SyncCompactorDeployments".
func (m *RisingWaveControllerManager) SyncCompactorDeployments() ctrlkit.Action {
	return ctrlkit.NewAction(RisingWaveAction_SyncCompactorDeployments, func(ctx context.Context) (result ctrl.Result, err error) {
		logger := m.logger.WithValues("action", RisingWaveAction_SyncCompactorDeployments)

		// Get states.
		compactorDeployments, err := m.state.GetCompactorDeployments(ctx)
		if err != nil {
			return ctrlkit.RequeueIfError(err)
		}

		// Invoke action.
		if m.hook != nil {
			defer func() { m.hook.PostRun(ctx, logger, RisingWaveAction_SyncCompactorDeployments, result, err) }()
			m.hook.PreRun(ctx, logger, RisingWaveAction_SyncCompactorDeployments, map[string]runtime.Object{
				"compactorDeployments": &appsv1.DeploymentList{Items: compactorDeployments},
			})
		}

		return m.impl.SyncCompactorDeployments(ctx, logger, compactorDeployments)
	})
}

// WaitBeforeCompactorDeploymentsReady generates the action of "WaitBeforeCompactorDeploymentsReady".
func (m *RisingWaveControllerManager) WaitBeforeCompactorDeploymentsReady() ctrlkit.Action {
	return ctrlkit.NewAction(RisingWaveAction_WaitBeforeCompactorDeploymentsReady, func(ctx context.Context) (result ctrl.Result, err error) {
		logger := m.logger.WithValues("action", RisingWaveAction_WaitBeforeCompactorDeploymentsReady)

		// Get states.
		compactorDeployments, err := m.state.GetCompactorDeployments(ctx)
		if err != nil {
			return ctrlkit.RequeueIfError(err)
		}

		// Invoke action.
		if m.hook != nil {
			defer func() { m.hook.PostRun(ctx, logger, RisingWaveAction_WaitBeforeCompactorDeploymentsReady, result, err) }()
			m.hook.PreRun(ctx, logger, RisingWaveAction_WaitBeforeCompactorDeploymentsReady, map[string]runtime.Object{
				"compactorDeployments": &appsv1.DeploymentList{Items: compactorDeployments},
			})
		}

		return m.impl.WaitBeforeCompactorDeploymentsReady(ctx, logger, compactorDeployments)
	})
}

// SyncConfigConfigMap generates the action of "SyncConfigConfigMap".
func (m *RisingWaveControllerManager) SyncConfigConfigMap() ctrlkit.Action {
	return ctrlkit.NewAction(RisingWaveAction_SyncConfigConfigMap, func(ctx context.Context) (result ctrl.Result, err error) {
		logger := m.logger.WithValues("action", RisingWaveAction_SyncConfigConfigMap)

		// Get states.
		configConfigMap, err := m.state.GetConfigConfigMap(ctx)
		if err != nil {
			return ctrlkit.RequeueIfError(err)
		}

		// Invoke action.
		if m.hook != nil {
			defer func() { m.hook.PostRun(ctx, logger, RisingWaveAction_SyncConfigConfigMap, result, err) }()
			m.hook.PreRun(ctx, logger, RisingWaveAction_SyncConfigConfigMap, map[string]runtime.Object{
				"configConfigMap": configConfigMap,
			})
		}

		return m.impl.SyncConfigConfigMap(ctx, logger, configConfigMap)
	})
}

// SyncServiceMonitor generates the action of "SyncServiceMonitor".
func (m *RisingWaveControllerManager) SyncServiceMonitor() ctrlkit.Action {
	return ctrlkit.NewAction(RisingWaveAction_SyncServiceMonitor, func(ctx context.Context) (result ctrl.Result, err error) {
		logger := m.logger.WithValues("action", RisingWaveAction_SyncServiceMonitor)

		// Get states.
		serviceMonitor, err := m.state.GetServiceMonitor(ctx)
		if err != nil {
			return ctrlkit.RequeueIfError(err)
		}

		// Invoke action.
		if m.hook != nil {
			defer func() { m.hook.PostRun(ctx, logger, RisingWaveAction_SyncServiceMonitor, result, err) }()
			m.hook.PreRun(ctx, logger, RisingWaveAction_SyncServiceMonitor, map[string]runtime.Object{
				"serviceMonitor": serviceMonitor,
			})
		}

		return m.impl.SyncServiceMonitor(ctx, logger, serviceMonitor)
	})
}

// CollectRunningStatisticsAndSyncStatus generates the action of "CollectRunningStatisticsAndSyncStatus".
func (m *RisingWaveControllerManager) CollectRunningStatisticsAndSyncStatus() ctrlkit.Action {
	return ctrlkit.NewAction(RisingWaveAction_CollectRunningStatisticsAndSyncStatus, func(ctx context.Context) (result ctrl.Result, err error) {
		logger := m.logger.WithValues("action", RisingWaveAction_CollectRunningStatisticsAndSyncStatus)

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

		metaDeployments, err := m.state.GetMetaDeployments(ctx)
		if err != nil {
			return ctrlkit.RequeueIfError(err)
		}

		frontendDeployments, err := m.state.GetFrontendDeployments(ctx)
		if err != nil {
			return ctrlkit.RequeueIfError(err)
		}

		computeStatefulSets, err := m.state.GetComputeStatefulSets(ctx)
		if err != nil {
			return ctrlkit.RequeueIfError(err)
		}

		compactorDeployments, err := m.state.GetCompactorDeployments(ctx)
		if err != nil {
			return ctrlkit.RequeueIfError(err)
		}

		configConfigMap, err := m.state.GetConfigConfigMap(ctx)
		if err != nil {
			return ctrlkit.RequeueIfError(err)
		}

		serviceMonitor, err := m.state.GetServiceMonitor(ctx)
		if err != nil {
			return ctrlkit.RequeueIfError(err)
		}

		// Invoke action.
		if m.hook != nil {
			defer func() {
				m.hook.PostRun(ctx, logger, RisingWaveAction_CollectRunningStatisticsAndSyncStatus, result, err)
			}()
			m.hook.PreRun(ctx, logger, RisingWaveAction_CollectRunningStatisticsAndSyncStatus, map[string]runtime.Object{
				"frontendService":      frontendService,
				"metaService":          metaService,
				"computeService":       computeService,
				"compactorService":     compactorService,
				"metaDeployments":      &appsv1.DeploymentList{Items: metaDeployments},
				"frontendDeployments":  &appsv1.DeploymentList{Items: frontendDeployments},
				"computeStatefulSets":  &appsv1.StatefulSetList{Items: computeStatefulSets},
				"compactorDeployments": &appsv1.DeploymentList{Items: compactorDeployments},
				"configConfigMap":      configConfigMap,
				"serviceMonitor":       serviceMonitor,
			})
		}

		return m.impl.CollectRunningStatisticsAndSyncStatus(ctx, logger, frontendService, metaService, computeService, compactorService, metaDeployments, frontendDeployments, computeStatefulSets, compactorDeployments, configConfigMap, serviceMonitor)
	})
}

type RisingWaveControllerManagerOption func(*RisingWaveControllerManager)

func RisingWaveControllerManager_WithActionHook(hook ctrlkit.ActionHook) RisingWaveControllerManagerOption {
	return func(m *RisingWaveControllerManager) {
		m.hook = hook
	}
}

// NewRisingWaveControllerManager returns a new RisingWaveControllerManager with given state and implementation.
func NewRisingWaveControllerManager(state RisingWaveControllerManagerState, impl RisingWaveControllerManagerImpl, logger logr.Logger, opts ...RisingWaveControllerManagerOption) RisingWaveControllerManager {
	m := RisingWaveControllerManager{
		state:  state,
		impl:   impl,
		logger: logger,
	}

	for _, opt := range opts {
		opt(&m)
	}

	return m
}
