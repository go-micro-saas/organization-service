package dbv1_0_0_org

import (
	"context"
	orgschemas "github.com/go-micro-saas/organization-service/app/org-service/internal/data/schema/org"
	employeeschemas "github.com/go-micro-saas/organization-service/app/org-service/internal/data/schema/org_employee"
	eventschemas "github.com/go-micro-saas/organization-service/app/org-service/internal/data/schema/org_event_history"
	inviteschemas "github.com/go-micro-saas/organization-service/app/org-service/internal/data/schema/org_invite_record"
	userorginfoschemas "github.com/go-micro-saas/organization-service/app/org-service/internal/data/schema/org_record_for_user"
	migrationpkg "github.com/ikaiguang/go-srv-kit/data/migration"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
	"gorm.io/gorm"
)

// Migrate 数据库迁移
type Migrate struct {
	dbConn      *gorm.DB
	migrateRepo migrationpkg.MigrateRepo
}

// NewMigrateHandler 处理手柄
func NewMigrateHandler(dbConn *gorm.DB, migrateRepo migrationpkg.MigrateRepo) *Migrate {
	return &Migrate{
		dbConn:      dbConn,
		migrateRepo: migrateRepo,
	}
}

// Upgrade ...
func (s *Migrate) Upgrade(ctx context.Context) error {
	var (
		mr       migrationpkg.MigrationInterface
		migrator = s.dbConn.WithContext(ctx).Migrator()
	)

	// 创建表
	mr = orgschemas.OrgSchema.CreateTableMigrator(migrator)
	if err := s.migrateRepo.RunMigratorUp(ctx, mr); err != nil {
		e := errorpkg.ErrorInternalError("")
		return errorpkg.Wrap(e, err)
	}
	mr = orgschemas.OrgSchema.AddColumnModifyStatusTime(migrator)
	if err := s.migrateRepo.RunMigratorUp(ctx, mr); err != nil {
		e := errorpkg.ErrorInternalError("")
		return errorpkg.Wrap(e, err)
	}
	// 创建表
	mr = employeeschemas.OrgEmployeeSchema.CreateTableMigrator(migrator)
	if err := s.migrateRepo.RunMigratorUp(ctx, mr); err != nil {
		e := errorpkg.ErrorInternalError("")
		return errorpkg.Wrap(e, err)
	}
	mr = employeeschemas.OrgEmployeeSchema.AddColumnModifyStatusTime(migrator)
	if err := s.migrateRepo.RunMigratorUp(ctx, mr); err != nil {
		e := errorpkg.ErrorInternalError("")
		return errorpkg.Wrap(e, err)
	}
	mr = employeeschemas.OrgEmployeeSchema.AddColumnModifyRoleTime(migrator)
	if err := s.migrateRepo.RunMigratorUp(ctx, mr); err != nil {
		e := errorpkg.ErrorInternalError("")
		return errorpkg.Wrap(e, err)
	}
	// 创建表
	mr = inviteschemas.OrgInviteRecordSchema.CreateTableMigrator(migrator)
	if err := s.migrateRepo.RunMigratorUp(ctx, mr); err != nil {
		e := errorpkg.ErrorInternalError("")
		return errorpkg.Wrap(e, err)
	}
	mr = userorginfoschemas.OrgRecordForUserSchema.CreateTableMigrator(migrator)
	if err := s.migrateRepo.RunMigratorUp(ctx, mr); err != nil {
		e := errorpkg.ErrorInternalError("")
		return errorpkg.Wrap(e, err)
	}
	// 创建表
	mr = eventschemas.OrgEventHistorySchema.CreateTableMigrator(migrator)
	if err := s.migrateRepo.RunMigratorUp(ctx, mr); err != nil {
		e := errorpkg.ErrorInternalError("")
		return errorpkg.Wrap(e, err)
	}
	return nil
}
