// Code generated from /grammar/MySqlParser.g4 by ANTLR 4.9.1. DO NOT EDIT.

package gen // MySqlParser
import "github.com/zeromicro/antlr"

type BaseMySqlParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseMySqlParserVisitor) VisitRoot(ctx *RootContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSqlStatements(ctx *SqlStatementsContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSqlStatement(ctx *SqlStatementContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitEmptyStatement(ctx *EmptyStatementContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitDdlStatement(ctx *DdlStatementContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitDmlStatement(ctx *DmlStatementContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTransactionStatement(ctx *TransactionStatementContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitReplicationStatement(ctx *ReplicationStatementContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitPreparedStatement(ctx *PreparedStatementContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitCompoundStatement(ctx *CompoundStatementContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAdministrationStatement(ctx *AdministrationStatementContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitUtilityStatement(ctx *UtilityStatementContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitCreateDatabase(ctx *CreateDatabaseContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitCreateEvent(ctx *CreateEventContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitCreateIndex(ctx *CreateIndexContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitCreateLogfileGroup(ctx *CreateLogfileGroupContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitCreateProcedure(ctx *CreateProcedureContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitCreateFunction(ctx *CreateFunctionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitCreateServer(ctx *CreateServerContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitCopyCreateTable(ctx *CopyCreateTableContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitQueryCreateTable(ctx *QueryCreateTableContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitColumnCreateTable(ctx *ColumnCreateTableContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitCreateTablespaceInnodb(ctx *CreateTablespaceInnodbContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitCreateTablespaceNdb(ctx *CreateTablespaceNdbContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitCreateTrigger(ctx *CreateTriggerContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitCreateView(ctx *CreateViewContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitCreateDatabaseOption(ctx *CreateDatabaseOptionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitOwnerStatement(ctx *OwnerStatementContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitPreciseSchedule(ctx *PreciseScheduleContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitIntervalSchedule(ctx *IntervalScheduleContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTimestampValue(ctx *TimestampValueContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitIntervalExpr(ctx *IntervalExprContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitIntervalType(ctx *IntervalTypeContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitEnableType(ctx *EnableTypeContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitIndexType(ctx *IndexTypeContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitIndexOption(ctx *IndexOptionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitProcedureParameter(ctx *ProcedureParameterContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitFunctionParameter(ctx *FunctionParameterContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitRoutineComment(ctx *RoutineCommentContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitRoutineLanguage(ctx *RoutineLanguageContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitRoutineBehavior(ctx *RoutineBehaviorContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitRoutineData(ctx *RoutineDataContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitRoutineSecurity(ctx *RoutineSecurityContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitServerOption(ctx *ServerOptionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitCreateDefinitions(ctx *CreateDefinitionsContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitColumnDeclaration(ctx *ColumnDeclarationContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitConstraintDeclaration(ctx *ConstraintDeclarationContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitIndexDeclaration(ctx *IndexDeclarationContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitColumnDefinition(ctx *ColumnDefinitionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitNullColumnConstraint(ctx *NullColumnConstraintContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitDefaultColumnConstraint(ctx *DefaultColumnConstraintContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAutoIncrementColumnConstraint(ctx *AutoIncrementColumnConstraintContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitPrimaryKeyColumnConstraint(ctx *PrimaryKeyColumnConstraintContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitUniqueKeyColumnConstraint(ctx *UniqueKeyColumnConstraintContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitCommentColumnConstraint(ctx *CommentColumnConstraintContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitFormatColumnConstraint(ctx *FormatColumnConstraintContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitStorageColumnConstraint(ctx *StorageColumnConstraintContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitReferenceColumnConstraint(ctx *ReferenceColumnConstraintContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitCollateColumnConstraint(ctx *CollateColumnConstraintContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitGeneratedColumnConstraint(ctx *GeneratedColumnConstraintContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSerialDefaultColumnConstraint(ctx *SerialDefaultColumnConstraintContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitCheckColumnConstraint(ctx *CheckColumnConstraintContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitPrimaryKeyTableConstraint(ctx *PrimaryKeyTableConstraintContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitUniqueKeyTableConstraint(ctx *UniqueKeyTableConstraintContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitForeignKeyTableConstraint(ctx *ForeignKeyTableConstraintContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitCheckTableConstraint(ctx *CheckTableConstraintContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitReferenceDefinition(ctx *ReferenceDefinitionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitReferenceAction(ctx *ReferenceActionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitReferenceControlType(ctx *ReferenceControlTypeContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSimpleIndexDeclaration(ctx *SimpleIndexDeclarationContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSpecialIndexDeclaration(ctx *SpecialIndexDeclarationContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTableOptionEngine(ctx *TableOptionEngineContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTableOptionAutoIncrement(ctx *TableOptionAutoIncrementContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTableOptionAverage(ctx *TableOptionAverageContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTableOptionCharset(ctx *TableOptionCharsetContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTableOptionChecksum(ctx *TableOptionChecksumContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTableOptionCollate(ctx *TableOptionCollateContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTableOptionComment(ctx *TableOptionCommentContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTableOptionCompression(ctx *TableOptionCompressionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTableOptionConnection(ctx *TableOptionConnectionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTableOptionDataDirectory(ctx *TableOptionDataDirectoryContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTableOptionDelay(ctx *TableOptionDelayContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTableOptionEncryption(ctx *TableOptionEncryptionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTableOptionIndexDirectory(ctx *TableOptionIndexDirectoryContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTableOptionInsertMethod(ctx *TableOptionInsertMethodContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTableOptionKeyBlockSize(ctx *TableOptionKeyBlockSizeContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTableOptionMaxRows(ctx *TableOptionMaxRowsContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTableOptionMinRows(ctx *TableOptionMinRowsContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTableOptionPackKeys(ctx *TableOptionPackKeysContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTableOptionPassword(ctx *TableOptionPasswordContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTableOptionRowFormat(ctx *TableOptionRowFormatContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTableOptionRecalculation(ctx *TableOptionRecalculationContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTableOptionPersistent(ctx *TableOptionPersistentContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTableOptionSamplePage(ctx *TableOptionSamplePageContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTableOptionTablespace(ctx *TableOptionTablespaceContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTableOptionUnion(ctx *TableOptionUnionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTablespaceStorage(ctx *TablespaceStorageContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitPartitionDefinitions(ctx *PartitionDefinitionsContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitPartitionFunctionHash(ctx *PartitionFunctionHashContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitPartitionFunctionKey(ctx *PartitionFunctionKeyContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitPartitionFunctionRange(ctx *PartitionFunctionRangeContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitPartitionFunctionList(ctx *PartitionFunctionListContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSubPartitionFunctionHash(ctx *SubPartitionFunctionHashContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSubPartitionFunctionKey(ctx *SubPartitionFunctionKeyContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitPartitionComparison(ctx *PartitionComparisonContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitPartitionListAtom(ctx *PartitionListAtomContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitPartitionListVector(ctx *PartitionListVectorContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitPartitionSimple(ctx *PartitionSimpleContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitPartitionDefinerAtom(ctx *PartitionDefinerAtomContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitPartitionDefinerVector(ctx *PartitionDefinerVectorContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSubpartitionDefinition(ctx *SubpartitionDefinitionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitPartitionOptionEngine(ctx *PartitionOptionEngineContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitPartitionOptionComment(ctx *PartitionOptionCommentContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitPartitionOptionDataDirectory(ctx *PartitionOptionDataDirectoryContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitPartitionOptionIndexDirectory(ctx *PartitionOptionIndexDirectoryContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitPartitionOptionMaxRows(ctx *PartitionOptionMaxRowsContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitPartitionOptionMinRows(ctx *PartitionOptionMinRowsContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitPartitionOptionTablespace(ctx *PartitionOptionTablespaceContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitPartitionOptionNodeGroup(ctx *PartitionOptionNodeGroupContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAlterSimpleDatabase(ctx *AlterSimpleDatabaseContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAlterUpgradeName(ctx *AlterUpgradeNameContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAlterEvent(ctx *AlterEventContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAlterFunction(ctx *AlterFunctionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAlterInstance(ctx *AlterInstanceContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAlterLogfileGroup(ctx *AlterLogfileGroupContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAlterProcedure(ctx *AlterProcedureContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAlterServer(ctx *AlterServerContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAlterTable(ctx *AlterTableContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAlterTablespace(ctx *AlterTablespaceContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAlterView(ctx *AlterViewContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAlterByTableOption(ctx *AlterByTableOptionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAlterByAddColumn(ctx *AlterByAddColumnContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAlterByAddColumns(ctx *AlterByAddColumnsContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAlterByAddIndex(ctx *AlterByAddIndexContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAlterByAddPrimaryKey(ctx *AlterByAddPrimaryKeyContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAlterByAddUniqueKey(ctx *AlterByAddUniqueKeyContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAlterByAddSpecialIndex(ctx *AlterByAddSpecialIndexContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAlterByAddForeignKey(ctx *AlterByAddForeignKeyContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAlterByAddCheckTableConstraint(ctx *AlterByAddCheckTableConstraintContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAlterBySetAlgorithm(ctx *AlterBySetAlgorithmContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAlterByChangeDefault(ctx *AlterByChangeDefaultContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAlterByChangeColumn(ctx *AlterByChangeColumnContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAlterByRenameColumn(ctx *AlterByRenameColumnContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAlterByLock(ctx *AlterByLockContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAlterByModifyColumn(ctx *AlterByModifyColumnContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAlterByDropColumn(ctx *AlterByDropColumnContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAlterByDropConstraintCheck(ctx *AlterByDropConstraintCheckContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAlterByDropPrimaryKey(ctx *AlterByDropPrimaryKeyContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAlterByRenameIndex(ctx *AlterByRenameIndexContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAlterByAlterIndexVisibility(ctx *AlterByAlterIndexVisibilityContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAlterByDropIndex(ctx *AlterByDropIndexContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAlterByDropForeignKey(ctx *AlterByDropForeignKeyContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAlterByDisableKeys(ctx *AlterByDisableKeysContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAlterByEnableKeys(ctx *AlterByEnableKeysContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAlterByRename(ctx *AlterByRenameContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAlterByOrder(ctx *AlterByOrderContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAlterByConvertCharset(ctx *AlterByConvertCharsetContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAlterByDefaultCharset(ctx *AlterByDefaultCharsetContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAlterByDiscardTablespace(ctx *AlterByDiscardTablespaceContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAlterByImportTablespace(ctx *AlterByImportTablespaceContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAlterByForce(ctx *AlterByForceContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAlterByValidate(ctx *AlterByValidateContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAlterByAddPartition(ctx *AlterByAddPartitionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAlterByDropPartition(ctx *AlterByDropPartitionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAlterByDiscardPartition(ctx *AlterByDiscardPartitionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAlterByImportPartition(ctx *AlterByImportPartitionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAlterByTruncatePartition(ctx *AlterByTruncatePartitionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAlterByCoalescePartition(ctx *AlterByCoalescePartitionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAlterByReorganizePartition(ctx *AlterByReorganizePartitionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAlterByExchangePartition(ctx *AlterByExchangePartitionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAlterByAnalyzePartition(ctx *AlterByAnalyzePartitionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAlterByCheckPartition(ctx *AlterByCheckPartitionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAlterByOptimizePartition(ctx *AlterByOptimizePartitionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAlterByRebuildPartition(ctx *AlterByRebuildPartitionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAlterByRepairPartition(ctx *AlterByRepairPartitionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAlterByRemovePartitioning(ctx *AlterByRemovePartitioningContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAlterByUpgradePartitioning(ctx *AlterByUpgradePartitioningContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitDropDatabase(ctx *DropDatabaseContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitDropEvent(ctx *DropEventContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitDropIndex(ctx *DropIndexContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitDropLogfileGroup(ctx *DropLogfileGroupContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitDropProcedure(ctx *DropProcedureContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitDropFunction(ctx *DropFunctionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitDropServer(ctx *DropServerContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitDropTable(ctx *DropTableContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitDropTablespace(ctx *DropTablespaceContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitDropTrigger(ctx *DropTriggerContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitDropView(ctx *DropViewContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitRenameTable(ctx *RenameTableContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitRenameTableClause(ctx *RenameTableClauseContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTruncateTable(ctx *TruncateTableContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitCallStatement(ctx *CallStatementContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitDeleteStatement(ctx *DeleteStatementContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitDoStatement(ctx *DoStatementContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitHandlerStatement(ctx *HandlerStatementContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitInsertStatement(ctx *InsertStatementContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitLoadDataStatement(ctx *LoadDataStatementContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitLoadXmlStatement(ctx *LoadXmlStatementContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitReplaceStatement(ctx *ReplaceStatementContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSimpleSelect(ctx *SimpleSelectContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitParenthesisSelect(ctx *ParenthesisSelectContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitUnionSelect(ctx *UnionSelectContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitUnionParenthesisSelect(ctx *UnionParenthesisSelectContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitUpdateStatement(ctx *UpdateStatementContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitInsertStatementValue(ctx *InsertStatementValueContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitUpdatedElement(ctx *UpdatedElementContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAssignmentField(ctx *AssignmentFieldContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitLockClause(ctx *LockClauseContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSingleDeleteStatement(ctx *SingleDeleteStatementContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitMultipleDeleteStatement(ctx *MultipleDeleteStatementContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitHandlerOpenStatement(ctx *HandlerOpenStatementContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitHandlerReadIndexStatement(ctx *HandlerReadIndexStatementContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitHandlerReadStatement(ctx *HandlerReadStatementContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitHandlerCloseStatement(ctx *HandlerCloseStatementContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSingleUpdateStatement(ctx *SingleUpdateStatementContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitMultipleUpdateStatement(ctx *MultipleUpdateStatementContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitOrderByClause(ctx *OrderByClauseContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitOrderByExpression(ctx *OrderByExpressionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTableSources(ctx *TableSourcesContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTableSourceBase(ctx *TableSourceBaseContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTableSourceNested(ctx *TableSourceNestedContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAtomTableItem(ctx *AtomTableItemContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSubqueryTableItem(ctx *SubqueryTableItemContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTableSourcesItem(ctx *TableSourcesItemContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitIndexHint(ctx *IndexHintContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitIndexHintType(ctx *IndexHintTypeContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitInnerJoin(ctx *InnerJoinContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitStraightJoin(ctx *StraightJoinContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitOuterJoin(ctx *OuterJoinContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitNaturalJoin(ctx *NaturalJoinContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitQueryExpression(ctx *QueryExpressionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitQueryExpressionNointo(ctx *QueryExpressionNointoContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitQuerySpecification(ctx *QuerySpecificationContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitQuerySpecificationNointo(ctx *QuerySpecificationNointoContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitUnionParenthesis(ctx *UnionParenthesisContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitUnionStatement(ctx *UnionStatementContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSelectSpec(ctx *SelectSpecContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSelectElements(ctx *SelectElementsContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSelectStarElement(ctx *SelectStarElementContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSelectColumnElement(ctx *SelectColumnElementContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSelectFunctionElement(ctx *SelectFunctionElementContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSelectExpressionElement(ctx *SelectExpressionElementContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSelectIntoVariables(ctx *SelectIntoVariablesContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSelectIntoDumpFile(ctx *SelectIntoDumpFileContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSelectIntoTextFile(ctx *SelectIntoTextFileContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSelectFieldsInto(ctx *SelectFieldsIntoContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSelectLinesInto(ctx *SelectLinesIntoContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitFromClause(ctx *FromClauseContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitGroupByClause(ctx *GroupByClauseContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitHavingClause(ctx *HavingClauseContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitGroupByItem(ctx *GroupByItemContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitLimitClause(ctx *LimitClauseContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitLimitClauseAtom(ctx *LimitClauseAtomContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitStartTransaction(ctx *StartTransactionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitBeginWork(ctx *BeginWorkContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitCommitWork(ctx *CommitWorkContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitRollbackWork(ctx *RollbackWorkContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSavepointStatement(ctx *SavepointStatementContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitRollbackStatement(ctx *RollbackStatementContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitReleaseStatement(ctx *ReleaseStatementContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitLockTables(ctx *LockTablesContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitUnlockTables(ctx *UnlockTablesContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSetAutocommitStatement(ctx *SetAutocommitStatementContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSetTransactionStatement(ctx *SetTransactionStatementContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTransactionMode(ctx *TransactionModeContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitLockTableElement(ctx *LockTableElementContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitLockAction(ctx *LockActionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTransactionOption(ctx *TransactionOptionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTransactionLevel(ctx *TransactionLevelContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitChangeMaster(ctx *ChangeMasterContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitChangeReplicationFilter(ctx *ChangeReplicationFilterContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitPurgeBinaryLogs(ctx *PurgeBinaryLogsContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitResetMaster(ctx *ResetMasterContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitResetSlave(ctx *ResetSlaveContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitStartSlave(ctx *StartSlaveContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitStopSlave(ctx *StopSlaveContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitStartGroupReplication(ctx *StartGroupReplicationContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitStopGroupReplication(ctx *StopGroupReplicationContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitMasterStringOption(ctx *MasterStringOptionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitMasterDecimalOption(ctx *MasterDecimalOptionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitMasterBoolOption(ctx *MasterBoolOptionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitMasterRealOption(ctx *MasterRealOptionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitMasterUidListOption(ctx *MasterUidListOptionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitStringMasterOption(ctx *StringMasterOptionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitDecimalMasterOption(ctx *DecimalMasterOptionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitBoolMasterOption(ctx *BoolMasterOptionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitChannelOption(ctx *ChannelOptionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitDoDbReplication(ctx *DoDbReplicationContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitIgnoreDbReplication(ctx *IgnoreDbReplicationContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitDoTableReplication(ctx *DoTableReplicationContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitIgnoreTableReplication(ctx *IgnoreTableReplicationContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitWildDoTableReplication(ctx *WildDoTableReplicationContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitWildIgnoreTableReplication(ctx *WildIgnoreTableReplicationContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitRewriteDbReplication(ctx *RewriteDbReplicationContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTablePair(ctx *TablePairContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitThreadType(ctx *ThreadTypeContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitGtidsUntilOption(ctx *GtidsUntilOptionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitMasterLogUntilOption(ctx *MasterLogUntilOptionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitRelayLogUntilOption(ctx *RelayLogUntilOptionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSqlGapsUntilOption(ctx *SqlGapsUntilOptionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitUserConnectionOption(ctx *UserConnectionOptionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitPasswordConnectionOption(ctx *PasswordConnectionOptionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitDefaultAuthConnectionOption(ctx *DefaultAuthConnectionOptionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitPluginDirConnectionOption(ctx *PluginDirConnectionOptionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitGtuidSet(ctx *GtuidSetContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitXaStartTransaction(ctx *XaStartTransactionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitXaEndTransaction(ctx *XaEndTransactionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitXaPrepareStatement(ctx *XaPrepareStatementContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitXaCommitWork(ctx *XaCommitWorkContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitXaRollbackWork(ctx *XaRollbackWorkContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitXaRecoverWork(ctx *XaRecoverWorkContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitPrepareStatement(ctx *PrepareStatementContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitExecuteStatement(ctx *ExecuteStatementContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitDeallocatePrepare(ctx *DeallocatePrepareContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitRoutineBody(ctx *RoutineBodyContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitBlockStatement(ctx *BlockStatementContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitCaseStatement(ctx *CaseStatementContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitIfStatement(ctx *IfStatementContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitIterateStatement(ctx *IterateStatementContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitLeaveStatement(ctx *LeaveStatementContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitLoopStatement(ctx *LoopStatementContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitRepeatStatement(ctx *RepeatStatementContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitReturnStatement(ctx *ReturnStatementContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitWhileStatement(ctx *WhileStatementContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitCloseCursor(ctx *CloseCursorContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitFetchCursor(ctx *FetchCursorContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitOpenCursor(ctx *OpenCursorContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitDeclareVariable(ctx *DeclareVariableContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitDeclareCondition(ctx *DeclareConditionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitDeclareCursor(ctx *DeclareCursorContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitDeclareHandler(ctx *DeclareHandlerContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitHandlerConditionCode(ctx *HandlerConditionCodeContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitHandlerConditionState(ctx *HandlerConditionStateContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitHandlerConditionName(ctx *HandlerConditionNameContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitHandlerConditionWarning(ctx *HandlerConditionWarningContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitHandlerConditionNotfound(ctx *HandlerConditionNotfoundContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitHandlerConditionException(ctx *HandlerConditionExceptionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitProcedureSqlStatement(ctx *ProcedureSqlStatementContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitCaseAlternative(ctx *CaseAlternativeContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitElifAlternative(ctx *ElifAlternativeContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAlterUserMysqlV56(ctx *AlterUserMysqlV56Context) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAlterUserMysqlV57(ctx *AlterUserMysqlV57Context) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitCreateUserMysqlV56(ctx *CreateUserMysqlV56Context) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitCreateUserMysqlV57(ctx *CreateUserMysqlV57Context) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitDropUser(ctx *DropUserContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitGrantStatement(ctx *GrantStatementContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitGrantProxy(ctx *GrantProxyContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitRenameUser(ctx *RenameUserContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitDetailRevoke(ctx *DetailRevokeContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitShortRevoke(ctx *ShortRevokeContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitRevokeProxy(ctx *RevokeProxyContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSetPasswordStatement(ctx *SetPasswordStatementContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitUserSpecification(ctx *UserSpecificationContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitPasswordAuthOption(ctx *PasswordAuthOptionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitStringAuthOption(ctx *StringAuthOptionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitHashAuthOption(ctx *HashAuthOptionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSimpleAuthOption(ctx *SimpleAuthOptionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTlsOption(ctx *TlsOptionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitUserResourceOption(ctx *UserResourceOptionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitUserPasswordOption(ctx *UserPasswordOptionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitUserLockOption(ctx *UserLockOptionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitPrivelegeClause(ctx *PrivelegeClauseContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitPrivilege(ctx *PrivilegeContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitCurrentSchemaPriviLevel(ctx *CurrentSchemaPriviLevelContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitGlobalPrivLevel(ctx *GlobalPrivLevelContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitDefiniteSchemaPrivLevel(ctx *DefiniteSchemaPrivLevelContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitDefiniteFullTablePrivLevel(ctx *DefiniteFullTablePrivLevelContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitDefiniteFullTablePrivLevel2(ctx *DefiniteFullTablePrivLevel2Context) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitDefiniteTablePrivLevel(ctx *DefiniteTablePrivLevelContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitRenameUserClause(ctx *RenameUserClauseContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAnalyzeTable(ctx *AnalyzeTableContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitCheckTable(ctx *CheckTableContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitChecksumTable(ctx *ChecksumTableContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitOptimizeTable(ctx *OptimizeTableContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitRepairTable(ctx *RepairTableContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitCheckTableOption(ctx *CheckTableOptionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitCreateUdfunction(ctx *CreateUdfunctionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitInstallPlugin(ctx *InstallPluginContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitUninstallPlugin(ctx *UninstallPluginContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSetVariable(ctx *SetVariableContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSetCharset(ctx *SetCharsetContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSetNames(ctx *SetNamesContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSetPassword(ctx *SetPasswordContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSetTransaction(ctx *SetTransactionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSetAutocommit(ctx *SetAutocommitContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSetNewValueInsideTrigger(ctx *SetNewValueInsideTriggerContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitShowMasterLogs(ctx *ShowMasterLogsContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitShowLogEvents(ctx *ShowLogEventsContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitShowObjectFilter(ctx *ShowObjectFilterContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitShowColumns(ctx *ShowColumnsContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitShowCreateDb(ctx *ShowCreateDbContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitShowCreateFullIdObject(ctx *ShowCreateFullIdObjectContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitShowCreateUser(ctx *ShowCreateUserContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitShowEngine(ctx *ShowEngineContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitShowGlobalInfo(ctx *ShowGlobalInfoContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitShowErrors(ctx *ShowErrorsContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitShowCountErrors(ctx *ShowCountErrorsContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitShowSchemaFilter(ctx *ShowSchemaFilterContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitShowRoutine(ctx *ShowRoutineContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitShowGrants(ctx *ShowGrantsContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitShowIndexes(ctx *ShowIndexesContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitShowOpenTables(ctx *ShowOpenTablesContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitShowProfile(ctx *ShowProfileContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitShowSlaveStatus(ctx *ShowSlaveStatusContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitVariableClause(ctx *VariableClauseContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitShowCommonEntity(ctx *ShowCommonEntityContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitShowFilter(ctx *ShowFilterContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitShowGlobalInfoClause(ctx *ShowGlobalInfoClauseContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitShowSchemaEntity(ctx *ShowSchemaEntityContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitShowProfileType(ctx *ShowProfileTypeContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitBinlogStatement(ctx *BinlogStatementContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitCacheIndexStatement(ctx *CacheIndexStatementContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitFlushStatement(ctx *FlushStatementContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitKillStatement(ctx *KillStatementContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitLoadIndexIntoCache(ctx *LoadIndexIntoCacheContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitResetStatement(ctx *ResetStatementContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitShutdownStatement(ctx *ShutdownStatementContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTableIndexes(ctx *TableIndexesContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSimpleFlushOption(ctx *SimpleFlushOptionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitChannelFlushOption(ctx *ChannelFlushOptionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTableFlushOption(ctx *TableFlushOptionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitFlushTableOption(ctx *FlushTableOptionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitLoadedTableIndexes(ctx *LoadedTableIndexesContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSimpleDescribeStatement(ctx *SimpleDescribeStatementContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitFullDescribeStatement(ctx *FullDescribeStatementContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitHelpStatement(ctx *HelpStatementContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitUseStatement(ctx *UseStatementContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSignalStatement(ctx *SignalStatementContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitResignalStatement(ctx *ResignalStatementContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSignalConditionInformation(ctx *SignalConditionInformationContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitDiagnosticsStatement(ctx *DiagnosticsStatementContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitDiagnosticsConditionInformationName(ctx *DiagnosticsConditionInformationNameContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitDescribeStatements(ctx *DescribeStatementsContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitDescribeConnection(ctx *DescribeConnectionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitFullId(ctx *FullIdContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTableName(ctx *TableNameContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitFullColumnName(ctx *FullColumnNameContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitIndexColumnName(ctx *IndexColumnNameContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitUserName(ctx *UserNameContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitMysqlVariable(ctx *MysqlVariableContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitCharsetName(ctx *CharsetNameContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitCollationName(ctx *CollationNameContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitEngineName(ctx *EngineNameContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitUuidSet(ctx *UuidSetContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitXid(ctx *XidContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitXuidStringId(ctx *XuidStringIdContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAuthPlugin(ctx *AuthPluginContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitUid(ctx *UidContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSimpleId(ctx *SimpleIdContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitDottedId(ctx *DottedIdContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitDecimalLiteral(ctx *DecimalLiteralContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitFileSizeLiteral(ctx *FileSizeLiteralContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitStringLiteral(ctx *StringLiteralContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitBooleanLiteral(ctx *BooleanLiteralContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitHexadecimalLiteral(ctx *HexadecimalLiteralContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitNullNotnull(ctx *NullNotnullContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitConstant(ctx *ConstantContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitStringDataType(ctx *StringDataTypeContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitNationalStringDataType(ctx *NationalStringDataTypeContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitNationalVaryingStringDataType(ctx *NationalVaryingStringDataTypeContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitDimensionDataType(ctx *DimensionDataTypeContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSimpleDataType(ctx *SimpleDataTypeContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitCollectionDataType(ctx *CollectionDataTypeContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSpatialDataType(ctx *SpatialDataTypeContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitLongVarcharDataType(ctx *LongVarcharDataTypeContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitLongVarbinaryDataType(ctx *LongVarbinaryDataTypeContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitCollectionOptions(ctx *CollectionOptionsContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitConvertedDataType(ctx *ConvertedDataTypeContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitLengthOneDimension(ctx *LengthOneDimensionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitLengthTwoDimension(ctx *LengthTwoDimensionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitLengthTwoOptionalDimension(ctx *LengthTwoOptionalDimensionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitUidList(ctx *UidListContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTables(ctx *TablesContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitIndexColumnNames(ctx *IndexColumnNamesContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitExpressions(ctx *ExpressionsContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitExpressionsWithDefaults(ctx *ExpressionsWithDefaultsContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitConstants(ctx *ConstantsContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSimpleStrings(ctx *SimpleStringsContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitUserVariables(ctx *UserVariablesContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitDefaultValue(ctx *DefaultValueContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitCurrentTimestamp(ctx *CurrentTimestampContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitExpressionOrDefault(ctx *ExpressionOrDefaultContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitIfExists(ctx *IfExistsContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitIfNotExists(ctx *IfNotExistsContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSpecificFunctionCall(ctx *SpecificFunctionCallContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAggregateFunctionCall(ctx *AggregateFunctionCallContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitScalarFunctionCall(ctx *ScalarFunctionCallContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitUdfFunctionCall(ctx *UdfFunctionCallContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitPasswordFunctionCall(ctx *PasswordFunctionCallContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSimpleFunctionCall(ctx *SimpleFunctionCallContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitDataTypeFunctionCall(ctx *DataTypeFunctionCallContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitValuesFunctionCall(ctx *ValuesFunctionCallContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitCaseExpressionFunctionCall(ctx *CaseExpressionFunctionCallContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitCaseFunctionCall(ctx *CaseFunctionCallContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitCharFunctionCall(ctx *CharFunctionCallContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitPositionFunctionCall(ctx *PositionFunctionCallContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSubstrFunctionCall(ctx *SubstrFunctionCallContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTrimFunctionCall(ctx *TrimFunctionCallContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitWeightFunctionCall(ctx *WeightFunctionCallContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitExtractFunctionCall(ctx *ExtractFunctionCallContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitGetFormatFunctionCall(ctx *GetFormatFunctionCallContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitJsonValueFunctionCall(ctx *JsonValueFunctionCallContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitCaseFuncAlternative(ctx *CaseFuncAlternativeContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitLevelWeightList(ctx *LevelWeightListContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitLevelWeightRange(ctx *LevelWeightRangeContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitLevelInWeightListElement(ctx *LevelInWeightListElementContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitAggregateWindowedFunction(ctx *AggregateWindowedFunctionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitScalarFunctionName(ctx *ScalarFunctionNameContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitPasswordFunctionClause(ctx *PasswordFunctionClauseContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitFunctionArgs(ctx *FunctionArgsContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitFunctionArg(ctx *FunctionArgContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitIsExpression(ctx *IsExpressionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitNotExpression(ctx *NotExpressionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitLogicalExpression(ctx *LogicalExpressionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitPredicateExpression(ctx *PredicateExpressionContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSoundsLikePredicate(ctx *SoundsLikePredicateContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitExpressionAtomPredicate(ctx *ExpressionAtomPredicateContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSubqueryComparisonPredicate(ctx *SubqueryComparisonPredicateContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitJsonMemberOfPredicate(ctx *JsonMemberOfPredicateContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitBinaryComparisonPredicate(ctx *BinaryComparisonPredicateContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitInPredicate(ctx *InPredicateContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitBetweenPredicate(ctx *BetweenPredicateContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitIsNullPredicate(ctx *IsNullPredicateContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitLikePredicate(ctx *LikePredicateContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitRegexpPredicate(ctx *RegexpPredicateContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitUnaryExpressionAtom(ctx *UnaryExpressionAtomContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitCollateExpressionAtom(ctx *CollateExpressionAtomContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitMysqlVariableExpressionAtom(ctx *MysqlVariableExpressionAtomContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitNestedExpressionAtom(ctx *NestedExpressionAtomContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitNestedRowExpressionAtom(ctx *NestedRowExpressionAtomContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitMathExpressionAtom(ctx *MathExpressionAtomContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitExistsExpressionAtom(ctx *ExistsExpressionAtomContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitIntervalExpressionAtom(ctx *IntervalExpressionAtomContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitJsonExpressionAtom(ctx *JsonExpressionAtomContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitSubqueryExpressionAtom(ctx *SubqueryExpressionAtomContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitConstantExpressionAtom(ctx *ConstantExpressionAtomContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitFunctionCallExpressionAtom(ctx *FunctionCallExpressionAtomContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitBinaryExpressionAtom(ctx *BinaryExpressionAtomContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitFullColumnNameExpressionAtom(ctx *FullColumnNameExpressionAtomContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitBitExpressionAtom(ctx *BitExpressionAtomContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitUnaryOperator(ctx *UnaryOperatorContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitComparisonOperator(ctx *ComparisonOperatorContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitLogicalOperator(ctx *LogicalOperatorContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitBitOperator(ctx *BitOperatorContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitMathOperator(ctx *MathOperatorContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitJsonOperator(ctx *JsonOperatorContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitCharsetNameBase(ctx *CharsetNameBaseContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitTransactionLevelBase(ctx *TransactionLevelBaseContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitPrivilegesBase(ctx *PrivilegesBaseContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitIntervalTypeBase(ctx *IntervalTypeBaseContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitDataTypeBase(ctx *DataTypeBaseContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitKeywordsCanBeId(ctx *KeywordsCanBeIdContext) any {
	return v.VisitChildren(ctx)
}

func (v *BaseMySqlParserVisitor) VisitFunctionNameBase(ctx *FunctionNameBaseContext) any {
	return v.VisitChildren(ctx)
}
