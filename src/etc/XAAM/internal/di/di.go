package di

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/syafdia/xaam/internal/domain/repo"
	au "github.com/syafdia/xaam/internal/domain/usecase/auth"
	bu "github.com/syafdia/xaam/internal/domain/usecase/business"
	"github.com/syafdia/xaam/internal/source"
)

type AppModule struct {
	DB *sqlx.DB
}

func GetAppModule() AppModule {
	dbConnection := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"))

	db, err := sqlx.Connect("postgres", dbConnection)
	if err != nil {
		log.Panic("failed initializing SQLX", err)
	}

	return AppModule{
		DB: db,
	}
}

type RepoModule struct {
	PolicyRepo   repo.PolicyRepo
	ResourceRepo repo.ResourceRepo
	BusinessRepo repo.BusinessRepo
}

func GetRepoModule(appModule AppModule) RepoModule {
	return RepoModule{
		PolicyRepo:   source.NewPolicyRepo(appModule.DB),
		ResourceRepo: source.NewResourceRepo(appModule.DB),
		BusinessRepo: source.NewBusinessRepo(),
	}
}

type UseCaseModule struct {
	FindResourcesByComplianceUseCase au.FindResourcesByComplianceUseCase
	FindOneByBusinessIDUseCase       bu.FindOneByBusinessIDUseCase
}

func GetUseCaseModule(repoModule RepoModule) UseCaseModule {
	return UseCaseModule{
		FindResourcesByComplianceUseCase: au.NewFindResourcesByComplianceUseCase(
			repoModule.PolicyRepo,
			repoModule.ResourceRepo,
		),
		FindOneByBusinessIDUseCase: bu.NewFindOneByBusinessIDUseCase(repoModule.BusinessRepo),
	}
}
