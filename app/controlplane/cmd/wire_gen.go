// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/chainloop-dev/chainloop/app/controlplane/internal/biz"
	"github.com/chainloop-dev/chainloop/app/controlplane/internal/conf"
	"github.com/chainloop-dev/chainloop/app/controlplane/internal/data"
	"github.com/chainloop-dev/chainloop/app/controlplane/internal/dispatcher"
	"github.com/chainloop-dev/chainloop/app/controlplane/internal/server"
	"github.com/chainloop-dev/chainloop/app/controlplane/internal/service"
	"github.com/chainloop-dev/chainloop/app/controlplane/plugins/sdk/v1"
	"github.com/chainloop-dev/chainloop/internal/blobmanager/loader"
	"github.com/chainloop-dev/chainloop/internal/credentials"
	"github.com/go-kratos/kratos/v2/log"
)

// Injectors from wire.go:

func wireApp(bootstrap *conf.Bootstrap, readerWriter credentials.ReaderWriter, logger log.Logger, availablePlugins sdk.AvailablePlugins) (*app, func(), error) {
	confData := bootstrap.Data
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	userRepo := data.NewUserRepo(dataData, logger)
	membershipRepo := data.NewMembershipRepo(dataData, logger)
	membershipUseCase := biz.NewMembershipUseCase(membershipRepo, logger)
	organizationRepo := data.NewOrganizationRepo(dataData, logger)
	casBackendRepo := data.NewCASBackendRepo(dataData, logger)
	providers := loader.LoadProviders(readerWriter)
	casBackendUseCase := biz.NewCASBackendUseCase(casBackendRepo, readerWriter, providers, logger)
	integrationRepo := data.NewIntegrationRepo(dataData, logger)
	integrationAttachmentRepo := data.NewIntegrationAttachmentRepo(dataData, logger)
	workflowRepo := data.NewWorkflowRepo(dataData, logger)
	newIntegrationUseCaseOpts := &biz.NewIntegrationUseCaseOpts{
		IRepo:   integrationRepo,
		IaRepo:  integrationAttachmentRepo,
		WfRepo:  workflowRepo,
		CredsRW: readerWriter,
		Logger:  logger,
	}
	integrationUseCase := biz.NewIntegrationUseCase(newIntegrationUseCaseOpts)
	organizationUseCase := biz.NewOrganizationUseCase(organizationRepo, casBackendUseCase, integrationUseCase, membershipRepo, logger)
	newUserUseCaseParams := &biz.NewUserUseCaseParams{
		UserRepo:            userRepo,
		MembershipUseCase:   membershipUseCase,
		OrganizationUseCase: organizationUseCase,
		Logger:              logger,
	}
	userUseCase := biz.NewUserUseCase(newUserUseCaseParams)
	robotAccountRepo := data.NewRobotAccountRepo(dataData, logger)
	auth := bootstrap.Auth
	robotAccountUseCase := biz.NewRootAccountUseCase(robotAccountRepo, workflowRepo, auth, logger)
	casCredentialsUseCase, err := biz.NewCASCredentialsUseCase(auth)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	bootstrap_CASServer := bootstrap.CasServer
	v := _wireValue
	casClientUseCase := biz.NewCASClientUseCase(casCredentialsUseCase, bootstrap_CASServer, logger, v...)
	referrerRepo := data.NewReferrerRepo(dataData, workflowRepo, logger)
	referrerSharedIndex := bootstrap.ReferrerSharedIndex
	referrerUseCase, err := biz.NewReferrerUseCase(referrerRepo, workflowRepo, membershipRepo, referrerSharedIndex, logger)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	workflowContractRepo := data.NewWorkflowContractRepo(dataData, logger)
	workflowContractUseCase := biz.NewWorkflowContractUseCase(workflowContractRepo, logger)
	workflowUseCase := biz.NewWorkflowUsecase(workflowRepo, workflowContractUseCase, logger)
	v2 := serviceOpts(logger)
	workflowService := service.NewWorkflowService(workflowUseCase, v2...)
	orgInvitationRepo := data.NewOrgInvitation(dataData, logger)
	orgInvitationUseCase, err := biz.NewOrgInvitationUseCase(orgInvitationRepo, membershipRepo, userRepo, logger)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	confServer := bootstrap.Server
	authService, err := service.NewAuthService(userUseCase, organizationUseCase, membershipUseCase, casBackendUseCase, orgInvitationUseCase, auth, confServer, v2...)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	robotAccountService := service.NewRobotAccountService(robotAccountUseCase, v2...)
	workflowRunRepo := data.NewWorkflowRunRepo(dataData, logger)
	workflowRunUseCase, err := biz.NewWorkflowRunUseCase(workflowRunRepo, workflowRepo, logger)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	newWorkflowRunServiceOpts := &service.NewWorkflowRunServiceOpts{
		WorkflowRunUC:      workflowRunUseCase,
		WorkflowUC:         workflowUseCase,
		WorkflowContractUC: workflowContractUseCase,
		CredsReader:        readerWriter,
		Opts:               v2,
	}
	workflowRunService := service.NewWorkflowRunService(newWorkflowRunServiceOpts)
	attestationUseCase := biz.NewAttestationUseCase(casClientUseCase, logger)
	fanOutDispatcher := dispatcher.New(integrationUseCase, workflowUseCase, workflowRunUseCase, readerWriter, casClientUseCase, availablePlugins, logger)
	casMappingRepo := data.NewCASMappingRepo(dataData, casBackendRepo, logger)
	casMappingUseCase := biz.NewCASMappingUseCase(casMappingRepo, membershipRepo, logger)
	newAttestationServiceOpts := &service.NewAttestationServiceOpts{
		WorkflowRunUC:      workflowRunUseCase,
		WorkflowUC:         workflowUseCase,
		WorkflowContractUC: workflowContractUseCase,
		OCIUC:              casBackendUseCase,
		CredsReader:        readerWriter,
		IntegrationUseCase: integrationUseCase,
		CasCredsUseCase:    casCredentialsUseCase,
		AttestationUC:      attestationUseCase,
		FanoutDispatcher:   fanOutDispatcher,
		CASMappingUseCase:  casMappingUseCase,
		ReferrerUC:         referrerUseCase,
		Opts:               v2,
	}
	attestationService := service.NewAttestationService(newAttestationServiceOpts)
	workflowContractService := service.NewWorkflowSchemaService(workflowContractUseCase, v2...)
	contextService := service.NewContextService(casBackendUseCase, v2...)
	casCredentialsService := service.NewCASCredentialsService(casCredentialsUseCase, casMappingUseCase, casBackendUseCase, v2...)
	orgMetricsRepo := data.NewOrgMetricsRepo(dataData, logger)
	orgMetricsUseCase, err := biz.NewOrgMetricsUseCase(orgMetricsRepo, logger)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	orgMetricsService := service.NewOrgMetricsService(orgMetricsUseCase, v2...)
	integrationsService := service.NewIntegrationsService(integrationUseCase, workflowUseCase, availablePlugins, v2...)
	organizationService := service.NewOrganizationService(membershipUseCase, organizationUseCase, v2...)
	casBackendService := service.NewCASBackendService(casBackendUseCase, providers, v2...)
	casRedirectService, err := service.NewCASRedirectService(casMappingUseCase, casCredentialsUseCase, bootstrap_CASServer, v2...)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	orgInvitationService := service.NewOrgInvitationService(orgInvitationUseCase, v2...)
	referrerService := service.NewReferrerService(referrerUseCase, v2...)
	opts := &server.Opts{
		UserUseCase:         userUseCase,
		RobotAccountUseCase: robotAccountUseCase,
		CASBackendUseCase:   casBackendUseCase,
		CASClientUseCase:    casClientUseCase,
		IntegrationUseCase:  integrationUseCase,
		ReferrerUseCase:     referrerUseCase,
		WorkflowSvc:         workflowService,
		AuthSvc:             authService,
		RobotAccountSvc:     robotAccountService,
		WorkflowRunSvc:      workflowRunService,
		AttestationSvc:      attestationService,
		WorkflowContractSvc: workflowContractService,
		ContextSvc:          contextService,
		CASCredsSvc:         casCredentialsService,
		OrgMetricsSvc:       orgMetricsService,
		IntegrationsSvc:     integrationsService,
		OrganizationSvc:     organizationService,
		CASBackendSvc:       casBackendService,
		CASRedirectSvc:      casRedirectService,
		OrgInvitationSvc:    orgInvitationService,
		ReferrerSvc:         referrerService,
		Logger:              logger,
		ServerConfig:        confServer,
		AuthConfig:          auth,
		Credentials:         readerWriter,
	}
	grpcServer, err := server.NewGRPCServer(opts)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	httpServer, err := server.NewHTTPServer(opts, grpcServer)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	httpMetricsServer, err := server.NewHTTPMetricsServer(opts)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	workflowRunExpirerUseCase := biz.NewWorkflowRunExpirerUseCase(workflowRunRepo, logger)
	mainApp := newApp(logger, grpcServer, httpServer, httpMetricsServer, workflowRunExpirerUseCase, availablePlugins)
	return mainApp, func() {
		cleanup()
	}, nil
}

var (
	_wireValue = []biz.CASClientOpts{}
)

// wire.go:

func serviceOpts(l log.Logger) []service.NewOpt {
	return []service.NewOpt{service.WithLogger(l)}
}
