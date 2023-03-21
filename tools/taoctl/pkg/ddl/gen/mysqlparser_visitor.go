// Code generated from /grammar/MySqlParser.g4 by ANTLR 4.9.1. DO NOT EDIT.

package gen // MySqlParser
import "github.com/zeromicro/antlr"

// A complete Visitor for a parse tree produced by MySqlParser.
type MySqlParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by MySqlParser#root.
	VisitRoot(ctx *RootContext) any

	// Visit a parse tree produced by MySqlParser#sqlStatements.
	VisitSqlStatements(ctx *SqlStatementsContext) any

	// Visit a parse tree produced by MySqlParser#sqlStatement.
	VisitSqlStatement(ctx *SqlStatementContext) any

	// Visit a parse tree produced by MySqlParser#emptyStatement.
	VisitEmptyStatement(ctx *EmptyStatementContext) any

	// Visit a parse tree produced by MySqlParser#ddlStatement.
	VisitDdlStatement(ctx *DdlStatementContext) any

	// Visit a parse tree produced by MySqlParser#dmlStatement.
	VisitDmlStatement(ctx *DmlStatementContext) any

	// Visit a parse tree produced by MySqlParser#transactionStatement.
	VisitTransactionStatement(ctx *TransactionStatementContext) any

	// Visit a parse tree produced by MySqlParser#replicationStatement.
	VisitReplicationStatement(ctx *ReplicationStatementContext) any

	// Visit a parse tree produced by MySqlParser#preparedStatement.
	VisitPreparedStatement(ctx *PreparedStatementContext) any

	// Visit a parse tree produced by MySqlParser#compoundStatement.
	VisitCompoundStatement(ctx *CompoundStatementContext) any

	// Visit a parse tree produced by MySqlParser#administrationStatement.
	VisitAdministrationStatement(ctx *AdministrationStatementContext) any

	// Visit a parse tree produced by MySqlParser#utilityStatement.
	VisitUtilityStatement(ctx *UtilityStatementContext) any

	// Visit a parse tree produced by MySqlParser#createDatabase.
	VisitCreateDatabase(ctx *CreateDatabaseContext) any

	// Visit a parse tree produced by MySqlParser#createEvent.
	VisitCreateEvent(ctx *CreateEventContext) any

	// Visit a parse tree produced by MySqlParser#createIndex.
	VisitCreateIndex(ctx *CreateIndexContext) any

	// Visit a parse tree produced by MySqlParser#createLogfileGroup.
	VisitCreateLogfileGroup(ctx *CreateLogfileGroupContext) any

	// Visit a parse tree produced by MySqlParser#createProcedure.
	VisitCreateProcedure(ctx *CreateProcedureContext) any

	// Visit a parse tree produced by MySqlParser#createFunction.
	VisitCreateFunction(ctx *CreateFunctionContext) any

	// Visit a parse tree produced by MySqlParser#createServer.
	VisitCreateServer(ctx *CreateServerContext) any

	// Visit a parse tree produced by MySqlParser#copyCreateTable.
	VisitCopyCreateTable(ctx *CopyCreateTableContext) any

	// Visit a parse tree produced by MySqlParser#queryCreateTable.
	VisitQueryCreateTable(ctx *QueryCreateTableContext) any

	// Visit a parse tree produced by MySqlParser#columnCreateTable.
	VisitColumnCreateTable(ctx *ColumnCreateTableContext) any

	// Visit a parse tree produced by MySqlParser#createTablespaceInnodb.
	VisitCreateTablespaceInnodb(ctx *CreateTablespaceInnodbContext) any

	// Visit a parse tree produced by MySqlParser#createTablespaceNdb.
	VisitCreateTablespaceNdb(ctx *CreateTablespaceNdbContext) any

	// Visit a parse tree produced by MySqlParser#createTrigger.
	VisitCreateTrigger(ctx *CreateTriggerContext) any

	// Visit a parse tree produced by MySqlParser#createView.
	VisitCreateView(ctx *CreateViewContext) any

	// Visit a parse tree produced by MySqlParser#createDatabaseOption.
	VisitCreateDatabaseOption(ctx *CreateDatabaseOptionContext) any

	// Visit a parse tree produced by MySqlParser#ownerStatement.
	VisitOwnerStatement(ctx *OwnerStatementContext) any

	// Visit a parse tree produced by MySqlParser#preciseSchedule.
	VisitPreciseSchedule(ctx *PreciseScheduleContext) any

	// Visit a parse tree produced by MySqlParser#intervalSchedule.
	VisitIntervalSchedule(ctx *IntervalScheduleContext) any

	// Visit a parse tree produced by MySqlParser#timestampValue.
	VisitTimestampValue(ctx *TimestampValueContext) any

	// Visit a parse tree produced by MySqlParser#intervalExpr.
	VisitIntervalExpr(ctx *IntervalExprContext) any

	// Visit a parse tree produced by MySqlParser#intervalType.
	VisitIntervalType(ctx *IntervalTypeContext) any

	// Visit a parse tree produced by MySqlParser#enableType.
	VisitEnableType(ctx *EnableTypeContext) any

	// Visit a parse tree produced by MySqlParser#indexType.
	VisitIndexType(ctx *IndexTypeContext) any

	// Visit a parse tree produced by MySqlParser#indexOption.
	VisitIndexOption(ctx *IndexOptionContext) any

	// Visit a parse tree produced by MySqlParser#procedureParameter.
	VisitProcedureParameter(ctx *ProcedureParameterContext) any

	// Visit a parse tree produced by MySqlParser#functionParameter.
	VisitFunctionParameter(ctx *FunctionParameterContext) any

	// Visit a parse tree produced by MySqlParser#routineComment.
	VisitRoutineComment(ctx *RoutineCommentContext) any

	// Visit a parse tree produced by MySqlParser#routineLanguage.
	VisitRoutineLanguage(ctx *RoutineLanguageContext) any

	// Visit a parse tree produced by MySqlParser#routineBehavior.
	VisitRoutineBehavior(ctx *RoutineBehaviorContext) any

	// Visit a parse tree produced by MySqlParser#routineData.
	VisitRoutineData(ctx *RoutineDataContext) any

	// Visit a parse tree produced by MySqlParser#routineSecurity.
	VisitRoutineSecurity(ctx *RoutineSecurityContext) any

	// Visit a parse tree produced by MySqlParser#serverOption.
	VisitServerOption(ctx *ServerOptionContext) any

	// Visit a parse tree produced by MySqlParser#createDefinitions.
	VisitCreateDefinitions(ctx *CreateDefinitionsContext) any

	// Visit a parse tree produced by MySqlParser#columnDeclaration.
	VisitColumnDeclaration(ctx *ColumnDeclarationContext) any

	// Visit a parse tree produced by MySqlParser#constraintDeclaration.
	VisitConstraintDeclaration(ctx *ConstraintDeclarationContext) any

	// Visit a parse tree produced by MySqlParser#indexDeclaration.
	VisitIndexDeclaration(ctx *IndexDeclarationContext) any

	// Visit a parse tree produced by MySqlParser#columnDefinition.
	VisitColumnDefinition(ctx *ColumnDefinitionContext) any

	// Visit a parse tree produced by MySqlParser#nullColumnConstraint.
	VisitNullColumnConstraint(ctx *NullColumnConstraintContext) any

	// Visit a parse tree produced by MySqlParser#defaultColumnConstraint.
	VisitDefaultColumnConstraint(ctx *DefaultColumnConstraintContext) any

	// Visit a parse tree produced by MySqlParser#autoIncrementColumnConstraint.
	VisitAutoIncrementColumnConstraint(ctx *AutoIncrementColumnConstraintContext) any

	// Visit a parse tree produced by MySqlParser#primaryKeyColumnConstraint.
	VisitPrimaryKeyColumnConstraint(ctx *PrimaryKeyColumnConstraintContext) any

	// Visit a parse tree produced by MySqlParser#uniqueKeyColumnConstraint.
	VisitUniqueKeyColumnConstraint(ctx *UniqueKeyColumnConstraintContext) any

	// Visit a parse tree produced by MySqlParser#commentColumnConstraint.
	VisitCommentColumnConstraint(ctx *CommentColumnConstraintContext) any

	// Visit a parse tree produced by MySqlParser#formatColumnConstraint.
	VisitFormatColumnConstraint(ctx *FormatColumnConstraintContext) any

	// Visit a parse tree produced by MySqlParser#storageColumnConstraint.
	VisitStorageColumnConstraint(ctx *StorageColumnConstraintContext) any

	// Visit a parse tree produced by MySqlParser#referenceColumnConstraint.
	VisitReferenceColumnConstraint(ctx *ReferenceColumnConstraintContext) any

	// Visit a parse tree produced by MySqlParser#collateColumnConstraint.
	VisitCollateColumnConstraint(ctx *CollateColumnConstraintContext) any

	// Visit a parse tree produced by MySqlParser#generatedColumnConstraint.
	VisitGeneratedColumnConstraint(ctx *GeneratedColumnConstraintContext) any

	// Visit a parse tree produced by MySqlParser#serialDefaultColumnConstraint.
	VisitSerialDefaultColumnConstraint(ctx *SerialDefaultColumnConstraintContext) any

	// Visit a parse tree produced by MySqlParser#checkColumnConstraint.
	VisitCheckColumnConstraint(ctx *CheckColumnConstraintContext) any

	// Visit a parse tree produced by MySqlParser#primaryKeyTableConstraint.
	VisitPrimaryKeyTableConstraint(ctx *PrimaryKeyTableConstraintContext) any

	// Visit a parse tree produced by MySqlParser#uniqueKeyTableConstraint.
	VisitUniqueKeyTableConstraint(ctx *UniqueKeyTableConstraintContext) any

	// Visit a parse tree produced by MySqlParser#foreignKeyTableConstraint.
	VisitForeignKeyTableConstraint(ctx *ForeignKeyTableConstraintContext) any

	// Visit a parse tree produced by MySqlParser#checkTableConstraint.
	VisitCheckTableConstraint(ctx *CheckTableConstraintContext) any

	// Visit a parse tree produced by MySqlParser#referenceDefinition.
	VisitReferenceDefinition(ctx *ReferenceDefinitionContext) any

	// Visit a parse tree produced by MySqlParser#referenceAction.
	VisitReferenceAction(ctx *ReferenceActionContext) any

	// Visit a parse tree produced by MySqlParser#referenceControlType.
	VisitReferenceControlType(ctx *ReferenceControlTypeContext) any

	// Visit a parse tree produced by MySqlParser#simpleIndexDeclaration.
	VisitSimpleIndexDeclaration(ctx *SimpleIndexDeclarationContext) any

	// Visit a parse tree produced by MySqlParser#specialIndexDeclaration.
	VisitSpecialIndexDeclaration(ctx *SpecialIndexDeclarationContext) any

	// Visit a parse tree produced by MySqlParser#tableOptionEngine.
	VisitTableOptionEngine(ctx *TableOptionEngineContext) any

	// Visit a parse tree produced by MySqlParser#tableOptionAutoIncrement.
	VisitTableOptionAutoIncrement(ctx *TableOptionAutoIncrementContext) any

	// Visit a parse tree produced by MySqlParser#tableOptionAverage.
	VisitTableOptionAverage(ctx *TableOptionAverageContext) any

	// Visit a parse tree produced by MySqlParser#tableOptionCharset.
	VisitTableOptionCharset(ctx *TableOptionCharsetContext) any

	// Visit a parse tree produced by MySqlParser#tableOptionChecksum.
	VisitTableOptionChecksum(ctx *TableOptionChecksumContext) any

	// Visit a parse tree produced by MySqlParser#tableOptionCollate.
	VisitTableOptionCollate(ctx *TableOptionCollateContext) any

	// Visit a parse tree produced by MySqlParser#tableOptionComment.
	VisitTableOptionComment(ctx *TableOptionCommentContext) any

	// Visit a parse tree produced by MySqlParser#tableOptionCompression.
	VisitTableOptionCompression(ctx *TableOptionCompressionContext) any

	// Visit a parse tree produced by MySqlParser#tableOptionConnection.
	VisitTableOptionConnection(ctx *TableOptionConnectionContext) any

	// Visit a parse tree produced by MySqlParser#tableOptionDataDirectory.
	VisitTableOptionDataDirectory(ctx *TableOptionDataDirectoryContext) any

	// Visit a parse tree produced by MySqlParser#tableOptionDelay.
	VisitTableOptionDelay(ctx *TableOptionDelayContext) any

	// Visit a parse tree produced by MySqlParser#tableOptionEncryption.
	VisitTableOptionEncryption(ctx *TableOptionEncryptionContext) any

	// Visit a parse tree produced by MySqlParser#tableOptionIndexDirectory.
	VisitTableOptionIndexDirectory(ctx *TableOptionIndexDirectoryContext) any

	// Visit a parse tree produced by MySqlParser#tableOptionInsertMethod.
	VisitTableOptionInsertMethod(ctx *TableOptionInsertMethodContext) any

	// Visit a parse tree produced by MySqlParser#tableOptionKeyBlockSize.
	VisitTableOptionKeyBlockSize(ctx *TableOptionKeyBlockSizeContext) any

	// Visit a parse tree produced by MySqlParser#tableOptionMaxRows.
	VisitTableOptionMaxRows(ctx *TableOptionMaxRowsContext) any

	// Visit a parse tree produced by MySqlParser#tableOptionMinRows.
	VisitTableOptionMinRows(ctx *TableOptionMinRowsContext) any

	// Visit a parse tree produced by MySqlParser#tableOptionPackKeys.
	VisitTableOptionPackKeys(ctx *TableOptionPackKeysContext) any

	// Visit a parse tree produced by MySqlParser#tableOptionPassword.
	VisitTableOptionPassword(ctx *TableOptionPasswordContext) any

	// Visit a parse tree produced by MySqlParser#tableOptionRowFormat.
	VisitTableOptionRowFormat(ctx *TableOptionRowFormatContext) any

	// Visit a parse tree produced by MySqlParser#tableOptionRecalculation.
	VisitTableOptionRecalculation(ctx *TableOptionRecalculationContext) any

	// Visit a parse tree produced by MySqlParser#tableOptionPersistent.
	VisitTableOptionPersistent(ctx *TableOptionPersistentContext) any

	// Visit a parse tree produced by MySqlParser#tableOptionSamplePage.
	VisitTableOptionSamplePage(ctx *TableOptionSamplePageContext) any

	// Visit a parse tree produced by MySqlParser#tableOptionTablespace.
	VisitTableOptionTablespace(ctx *TableOptionTablespaceContext) any

	// Visit a parse tree produced by MySqlParser#tableOptionUnion.
	VisitTableOptionUnion(ctx *TableOptionUnionContext) any

	// Visit a parse tree produced by MySqlParser#tablespaceStorage.
	VisitTablespaceStorage(ctx *TablespaceStorageContext) any

	// Visit a parse tree produced by MySqlParser#partitionDefinitions.
	VisitPartitionDefinitions(ctx *PartitionDefinitionsContext) any

	// Visit a parse tree produced by MySqlParser#partitionFunctionHash.
	VisitPartitionFunctionHash(ctx *PartitionFunctionHashContext) any

	// Visit a parse tree produced by MySqlParser#partitionFunctionKey.
	VisitPartitionFunctionKey(ctx *PartitionFunctionKeyContext) any

	// Visit a parse tree produced by MySqlParser#partitionFunctionRange.
	VisitPartitionFunctionRange(ctx *PartitionFunctionRangeContext) any

	// Visit a parse tree produced by MySqlParser#partitionFunctionList.
	VisitPartitionFunctionList(ctx *PartitionFunctionListContext) any

	// Visit a parse tree produced by MySqlParser#subPartitionFunctionHash.
	VisitSubPartitionFunctionHash(ctx *SubPartitionFunctionHashContext) any

	// Visit a parse tree produced by MySqlParser#subPartitionFunctionKey.
	VisitSubPartitionFunctionKey(ctx *SubPartitionFunctionKeyContext) any

	// Visit a parse tree produced by MySqlParser#partitionComparison.
	VisitPartitionComparison(ctx *PartitionComparisonContext) any

	// Visit a parse tree produced by MySqlParser#partitionListAtom.
	VisitPartitionListAtom(ctx *PartitionListAtomContext) any

	// Visit a parse tree produced by MySqlParser#partitionListVector.
	VisitPartitionListVector(ctx *PartitionListVectorContext) any

	// Visit a parse tree produced by MySqlParser#partitionSimple.
	VisitPartitionSimple(ctx *PartitionSimpleContext) any

	// Visit a parse tree produced by MySqlParser#partitionDefinerAtom.
	VisitPartitionDefinerAtom(ctx *PartitionDefinerAtomContext) any

	// Visit a parse tree produced by MySqlParser#partitionDefinerVector.
	VisitPartitionDefinerVector(ctx *PartitionDefinerVectorContext) any

	// Visit a parse tree produced by MySqlParser#subpartitionDefinition.
	VisitSubpartitionDefinition(ctx *SubpartitionDefinitionContext) any

	// Visit a parse tree produced by MySqlParser#partitionOptionEngine.
	VisitPartitionOptionEngine(ctx *PartitionOptionEngineContext) any

	// Visit a parse tree produced by MySqlParser#partitionOptionComment.
	VisitPartitionOptionComment(ctx *PartitionOptionCommentContext) any

	// Visit a parse tree produced by MySqlParser#partitionOptionDataDirectory.
	VisitPartitionOptionDataDirectory(ctx *PartitionOptionDataDirectoryContext) any

	// Visit a parse tree produced by MySqlParser#partitionOptionIndexDirectory.
	VisitPartitionOptionIndexDirectory(ctx *PartitionOptionIndexDirectoryContext) any

	// Visit a parse tree produced by MySqlParser#partitionOptionMaxRows.
	VisitPartitionOptionMaxRows(ctx *PartitionOptionMaxRowsContext) any

	// Visit a parse tree produced by MySqlParser#partitionOptionMinRows.
	VisitPartitionOptionMinRows(ctx *PartitionOptionMinRowsContext) any

	// Visit a parse tree produced by MySqlParser#partitionOptionTablespace.
	VisitPartitionOptionTablespace(ctx *PartitionOptionTablespaceContext) any

	// Visit a parse tree produced by MySqlParser#partitionOptionNodeGroup.
	VisitPartitionOptionNodeGroup(ctx *PartitionOptionNodeGroupContext) any

	// Visit a parse tree produced by MySqlParser#alterSimpleDatabase.
	VisitAlterSimpleDatabase(ctx *AlterSimpleDatabaseContext) any

	// Visit a parse tree produced by MySqlParser#alterUpgradeName.
	VisitAlterUpgradeName(ctx *AlterUpgradeNameContext) any

	// Visit a parse tree produced by MySqlParser#alterEvent.
	VisitAlterEvent(ctx *AlterEventContext) any

	// Visit a parse tree produced by MySqlParser#alterFunction.
	VisitAlterFunction(ctx *AlterFunctionContext) any

	// Visit a parse tree produced by MySqlParser#alterInstance.
	VisitAlterInstance(ctx *AlterInstanceContext) any

	// Visit a parse tree produced by MySqlParser#alterLogfileGroup.
	VisitAlterLogfileGroup(ctx *AlterLogfileGroupContext) any

	// Visit a parse tree produced by MySqlParser#alterProcedure.
	VisitAlterProcedure(ctx *AlterProcedureContext) any

	// Visit a parse tree produced by MySqlParser#alterServer.
	VisitAlterServer(ctx *AlterServerContext) any

	// Visit a parse tree produced by MySqlParser#alterTable.
	VisitAlterTable(ctx *AlterTableContext) any

	// Visit a parse tree produced by MySqlParser#alterTablespace.
	VisitAlterTablespace(ctx *AlterTablespaceContext) any

	// Visit a parse tree produced by MySqlParser#alterView.
	VisitAlterView(ctx *AlterViewContext) any

	// Visit a parse tree produced by MySqlParser#alterByTableOption.
	VisitAlterByTableOption(ctx *AlterByTableOptionContext) any

	// Visit a parse tree produced by MySqlParser#alterByAddColumn.
	VisitAlterByAddColumn(ctx *AlterByAddColumnContext) any

	// Visit a parse tree produced by MySqlParser#alterByAddColumns.
	VisitAlterByAddColumns(ctx *AlterByAddColumnsContext) any

	// Visit a parse tree produced by MySqlParser#alterByAddIndex.
	VisitAlterByAddIndex(ctx *AlterByAddIndexContext) any

	// Visit a parse tree produced by MySqlParser#alterByAddPrimaryKey.
	VisitAlterByAddPrimaryKey(ctx *AlterByAddPrimaryKeyContext) any

	// Visit a parse tree produced by MySqlParser#alterByAddUniqueKey.
	VisitAlterByAddUniqueKey(ctx *AlterByAddUniqueKeyContext) any

	// Visit a parse tree produced by MySqlParser#alterByAddSpecialIndex.
	VisitAlterByAddSpecialIndex(ctx *AlterByAddSpecialIndexContext) any

	// Visit a parse tree produced by MySqlParser#alterByAddForeignKey.
	VisitAlterByAddForeignKey(ctx *AlterByAddForeignKeyContext) any

	// Visit a parse tree produced by MySqlParser#alterByAddCheckTableConstraint.
	VisitAlterByAddCheckTableConstraint(ctx *AlterByAddCheckTableConstraintContext) any

	// Visit a parse tree produced by MySqlParser#alterBySetAlgorithm.
	VisitAlterBySetAlgorithm(ctx *AlterBySetAlgorithmContext) any

	// Visit a parse tree produced by MySqlParser#alterByChangeDefault.
	VisitAlterByChangeDefault(ctx *AlterByChangeDefaultContext) any

	// Visit a parse tree produced by MySqlParser#alterByChangeColumn.
	VisitAlterByChangeColumn(ctx *AlterByChangeColumnContext) any

	// Visit a parse tree produced by MySqlParser#alterByRenameColumn.
	VisitAlterByRenameColumn(ctx *AlterByRenameColumnContext) any

	// Visit a parse tree produced by MySqlParser#alterByLock.
	VisitAlterByLock(ctx *AlterByLockContext) any

	// Visit a parse tree produced by MySqlParser#alterByModifyColumn.
	VisitAlterByModifyColumn(ctx *AlterByModifyColumnContext) any

	// Visit a parse tree produced by MySqlParser#alterByDropColumn.
	VisitAlterByDropColumn(ctx *AlterByDropColumnContext) any

	// Visit a parse tree produced by MySqlParser#alterByDropConstraintCheck.
	VisitAlterByDropConstraintCheck(ctx *AlterByDropConstraintCheckContext) any

	// Visit a parse tree produced by MySqlParser#alterByDropPrimaryKey.
	VisitAlterByDropPrimaryKey(ctx *AlterByDropPrimaryKeyContext) any

	// Visit a parse tree produced by MySqlParser#alterByRenameIndex.
	VisitAlterByRenameIndex(ctx *AlterByRenameIndexContext) any

	// Visit a parse tree produced by MySqlParser#alterByAlterIndexVisibility.
	VisitAlterByAlterIndexVisibility(ctx *AlterByAlterIndexVisibilityContext) any

	// Visit a parse tree produced by MySqlParser#alterByDropIndex.
	VisitAlterByDropIndex(ctx *AlterByDropIndexContext) any

	// Visit a parse tree produced by MySqlParser#alterByDropForeignKey.
	VisitAlterByDropForeignKey(ctx *AlterByDropForeignKeyContext) any

	// Visit a parse tree produced by MySqlParser#alterByDisableKeys.
	VisitAlterByDisableKeys(ctx *AlterByDisableKeysContext) any

	// Visit a parse tree produced by MySqlParser#alterByEnableKeys.
	VisitAlterByEnableKeys(ctx *AlterByEnableKeysContext) any

	// Visit a parse tree produced by MySqlParser#alterByRename.
	VisitAlterByRename(ctx *AlterByRenameContext) any

	// Visit a parse tree produced by MySqlParser#alterByOrder.
	VisitAlterByOrder(ctx *AlterByOrderContext) any

	// Visit a parse tree produced by MySqlParser#alterByConvertCharset.
	VisitAlterByConvertCharset(ctx *AlterByConvertCharsetContext) any

	// Visit a parse tree produced by MySqlParser#alterByDefaultCharset.
	VisitAlterByDefaultCharset(ctx *AlterByDefaultCharsetContext) any

	// Visit a parse tree produced by MySqlParser#alterByDiscardTablespace.
	VisitAlterByDiscardTablespace(ctx *AlterByDiscardTablespaceContext) any

	// Visit a parse tree produced by MySqlParser#alterByImportTablespace.
	VisitAlterByImportTablespace(ctx *AlterByImportTablespaceContext) any

	// Visit a parse tree produced by MySqlParser#alterByForce.
	VisitAlterByForce(ctx *AlterByForceContext) any

	// Visit a parse tree produced by MySqlParser#alterByValidate.
	VisitAlterByValidate(ctx *AlterByValidateContext) any

	// Visit a parse tree produced by MySqlParser#alterByAddPartition.
	VisitAlterByAddPartition(ctx *AlterByAddPartitionContext) any

	// Visit a parse tree produced by MySqlParser#alterByDropPartition.
	VisitAlterByDropPartition(ctx *AlterByDropPartitionContext) any

	// Visit a parse tree produced by MySqlParser#alterByDiscardPartition.
	VisitAlterByDiscardPartition(ctx *AlterByDiscardPartitionContext) any

	// Visit a parse tree produced by MySqlParser#alterByImportPartition.
	VisitAlterByImportPartition(ctx *AlterByImportPartitionContext) any

	// Visit a parse tree produced by MySqlParser#alterByTruncatePartition.
	VisitAlterByTruncatePartition(ctx *AlterByTruncatePartitionContext) any

	// Visit a parse tree produced by MySqlParser#alterByCoalescePartition.
	VisitAlterByCoalescePartition(ctx *AlterByCoalescePartitionContext) any

	// Visit a parse tree produced by MySqlParser#alterByReorganizePartition.
	VisitAlterByReorganizePartition(ctx *AlterByReorganizePartitionContext) any

	// Visit a parse tree produced by MySqlParser#alterByExchangePartition.
	VisitAlterByExchangePartition(ctx *AlterByExchangePartitionContext) any

	// Visit a parse tree produced by MySqlParser#alterByAnalyzePartition.
	VisitAlterByAnalyzePartition(ctx *AlterByAnalyzePartitionContext) any

	// Visit a parse tree produced by MySqlParser#alterByCheckPartition.
	VisitAlterByCheckPartition(ctx *AlterByCheckPartitionContext) any

	// Visit a parse tree produced by MySqlParser#alterByOptimizePartition.
	VisitAlterByOptimizePartition(ctx *AlterByOptimizePartitionContext) any

	// Visit a parse tree produced by MySqlParser#alterByRebuildPartition.
	VisitAlterByRebuildPartition(ctx *AlterByRebuildPartitionContext) any

	// Visit a parse tree produced by MySqlParser#alterByRepairPartition.
	VisitAlterByRepairPartition(ctx *AlterByRepairPartitionContext) any

	// Visit a parse tree produced by MySqlParser#alterByRemovePartitioning.
	VisitAlterByRemovePartitioning(ctx *AlterByRemovePartitioningContext) any

	// Visit a parse tree produced by MySqlParser#alterByUpgradePartitioning.
	VisitAlterByUpgradePartitioning(ctx *AlterByUpgradePartitioningContext) any

	// Visit a parse tree produced by MySqlParser#dropDatabase.
	VisitDropDatabase(ctx *DropDatabaseContext) any

	// Visit a parse tree produced by MySqlParser#dropEvent.
	VisitDropEvent(ctx *DropEventContext) any

	// Visit a parse tree produced by MySqlParser#dropIndex.
	VisitDropIndex(ctx *DropIndexContext) any

	// Visit a parse tree produced by MySqlParser#dropLogfileGroup.
	VisitDropLogfileGroup(ctx *DropLogfileGroupContext) any

	// Visit a parse tree produced by MySqlParser#dropProcedure.
	VisitDropProcedure(ctx *DropProcedureContext) any

	// Visit a parse tree produced by MySqlParser#dropFunction.
	VisitDropFunction(ctx *DropFunctionContext) any

	// Visit a parse tree produced by MySqlParser#dropServer.
	VisitDropServer(ctx *DropServerContext) any

	// Visit a parse tree produced by MySqlParser#dropTable.
	VisitDropTable(ctx *DropTableContext) any

	// Visit a parse tree produced by MySqlParser#dropTablespace.
	VisitDropTablespace(ctx *DropTablespaceContext) any

	// Visit a parse tree produced by MySqlParser#dropTrigger.
	VisitDropTrigger(ctx *DropTriggerContext) any

	// Visit a parse tree produced by MySqlParser#dropView.
	VisitDropView(ctx *DropViewContext) any

	// Visit a parse tree produced by MySqlParser#renameTable.
	VisitRenameTable(ctx *RenameTableContext) any

	// Visit a parse tree produced by MySqlParser#renameTableClause.
	VisitRenameTableClause(ctx *RenameTableClauseContext) any

	// Visit a parse tree produced by MySqlParser#truncateTable.
	VisitTruncateTable(ctx *TruncateTableContext) any

	// Visit a parse tree produced by MySqlParser#callStatement.
	VisitCallStatement(ctx *CallStatementContext) any

	// Visit a parse tree produced by MySqlParser#deleteStatement.
	VisitDeleteStatement(ctx *DeleteStatementContext) any

	// Visit a parse tree produced by MySqlParser#doStatement.
	VisitDoStatement(ctx *DoStatementContext) any

	// Visit a parse tree produced by MySqlParser#handlerStatement.
	VisitHandlerStatement(ctx *HandlerStatementContext) any

	// Visit a parse tree produced by MySqlParser#insertStatement.
	VisitInsertStatement(ctx *InsertStatementContext) any

	// Visit a parse tree produced by MySqlParser#loadDataStatement.
	VisitLoadDataStatement(ctx *LoadDataStatementContext) any

	// Visit a parse tree produced by MySqlParser#loadXmlStatement.
	VisitLoadXmlStatement(ctx *LoadXmlStatementContext) any

	// Visit a parse tree produced by MySqlParser#replaceStatement.
	VisitReplaceStatement(ctx *ReplaceStatementContext) any

	// Visit a parse tree produced by MySqlParser#simpleSelect.
	VisitSimpleSelect(ctx *SimpleSelectContext) any

	// Visit a parse tree produced by MySqlParser#parenthesisSelect.
	VisitParenthesisSelect(ctx *ParenthesisSelectContext) any

	// Visit a parse tree produced by MySqlParser#unionSelect.
	VisitUnionSelect(ctx *UnionSelectContext) any

	// Visit a parse tree produced by MySqlParser#unionParenthesisSelect.
	VisitUnionParenthesisSelect(ctx *UnionParenthesisSelectContext) any

	// Visit a parse tree produced by MySqlParser#updateStatement.
	VisitUpdateStatement(ctx *UpdateStatementContext) any

	// Visit a parse tree produced by MySqlParser#insertStatementValue.
	VisitInsertStatementValue(ctx *InsertStatementValueContext) any

	// Visit a parse tree produced by MySqlParser#updatedElement.
	VisitUpdatedElement(ctx *UpdatedElementContext) any

	// Visit a parse tree produced by MySqlParser#assignmentField.
	VisitAssignmentField(ctx *AssignmentFieldContext) any

	// Visit a parse tree produced by MySqlParser#lockClause.
	VisitLockClause(ctx *LockClauseContext) any

	// Visit a parse tree produced by MySqlParser#singleDeleteStatement.
	VisitSingleDeleteStatement(ctx *SingleDeleteStatementContext) any

	// Visit a parse tree produced by MySqlParser#multipleDeleteStatement.
	VisitMultipleDeleteStatement(ctx *MultipleDeleteStatementContext) any

	// Visit a parse tree produced by MySqlParser#handlerOpenStatement.
	VisitHandlerOpenStatement(ctx *HandlerOpenStatementContext) any

	// Visit a parse tree produced by MySqlParser#handlerReadIndexStatement.
	VisitHandlerReadIndexStatement(ctx *HandlerReadIndexStatementContext) any

	// Visit a parse tree produced by MySqlParser#handlerReadStatement.
	VisitHandlerReadStatement(ctx *HandlerReadStatementContext) any

	// Visit a parse tree produced by MySqlParser#handlerCloseStatement.
	VisitHandlerCloseStatement(ctx *HandlerCloseStatementContext) any

	// Visit a parse tree produced by MySqlParser#singleUpdateStatement.
	VisitSingleUpdateStatement(ctx *SingleUpdateStatementContext) any

	// Visit a parse tree produced by MySqlParser#multipleUpdateStatement.
	VisitMultipleUpdateStatement(ctx *MultipleUpdateStatementContext) any

	// Visit a parse tree produced by MySqlParser#orderByClause.
	VisitOrderByClause(ctx *OrderByClauseContext) any

	// Visit a parse tree produced by MySqlParser#orderByExpression.
	VisitOrderByExpression(ctx *OrderByExpressionContext) any

	// Visit a parse tree produced by MySqlParser#tableSources.
	VisitTableSources(ctx *TableSourcesContext) any

	// Visit a parse tree produced by MySqlParser#tableSourceBase.
	VisitTableSourceBase(ctx *TableSourceBaseContext) any

	// Visit a parse tree produced by MySqlParser#tableSourceNested.
	VisitTableSourceNested(ctx *TableSourceNestedContext) any

	// Visit a parse tree produced by MySqlParser#atomTableItem.
	VisitAtomTableItem(ctx *AtomTableItemContext) any

	// Visit a parse tree produced by MySqlParser#subqueryTableItem.
	VisitSubqueryTableItem(ctx *SubqueryTableItemContext) any

	// Visit a parse tree produced by MySqlParser#tableSourcesItem.
	VisitTableSourcesItem(ctx *TableSourcesItemContext) any

	// Visit a parse tree produced by MySqlParser#indexHint.
	VisitIndexHint(ctx *IndexHintContext) any

	// Visit a parse tree produced by MySqlParser#indexHintType.
	VisitIndexHintType(ctx *IndexHintTypeContext) any

	// Visit a parse tree produced by MySqlParser#innerJoin.
	VisitInnerJoin(ctx *InnerJoinContext) any

	// Visit a parse tree produced by MySqlParser#straightJoin.
	VisitStraightJoin(ctx *StraightJoinContext) any

	// Visit a parse tree produced by MySqlParser#outerJoin.
	VisitOuterJoin(ctx *OuterJoinContext) any

	// Visit a parse tree produced by MySqlParser#naturalJoin.
	VisitNaturalJoin(ctx *NaturalJoinContext) any

	// Visit a parse tree produced by MySqlParser#queryExpression.
	VisitQueryExpression(ctx *QueryExpressionContext) any

	// Visit a parse tree produced by MySqlParser#queryExpressionNointo.
	VisitQueryExpressionNointo(ctx *QueryExpressionNointoContext) any

	// Visit a parse tree produced by MySqlParser#querySpecification.
	VisitQuerySpecification(ctx *QuerySpecificationContext) any

	// Visit a parse tree produced by MySqlParser#querySpecificationNointo.
	VisitQuerySpecificationNointo(ctx *QuerySpecificationNointoContext) any

	// Visit a parse tree produced by MySqlParser#unionParenthesis.
	VisitUnionParenthesis(ctx *UnionParenthesisContext) any

	// Visit a parse tree produced by MySqlParser#unionStatement.
	VisitUnionStatement(ctx *UnionStatementContext) any

	// Visit a parse tree produced by MySqlParser#selectSpec.
	VisitSelectSpec(ctx *SelectSpecContext) any

	// Visit a parse tree produced by MySqlParser#selectElements.
	VisitSelectElements(ctx *SelectElementsContext) any

	// Visit a parse tree produced by MySqlParser#selectStarElement.
	VisitSelectStarElement(ctx *SelectStarElementContext) any

	// Visit a parse tree produced by MySqlParser#selectColumnElement.
	VisitSelectColumnElement(ctx *SelectColumnElementContext) any

	// Visit a parse tree produced by MySqlParser#selectFunctionElement.
	VisitSelectFunctionElement(ctx *SelectFunctionElementContext) any

	// Visit a parse tree produced by MySqlParser#selectExpressionElement.
	VisitSelectExpressionElement(ctx *SelectExpressionElementContext) any

	// Visit a parse tree produced by MySqlParser#selectIntoVariables.
	VisitSelectIntoVariables(ctx *SelectIntoVariablesContext) any

	// Visit a parse tree produced by MySqlParser#selectIntoDumpFile.
	VisitSelectIntoDumpFile(ctx *SelectIntoDumpFileContext) any

	// Visit a parse tree produced by MySqlParser#selectIntoTextFile.
	VisitSelectIntoTextFile(ctx *SelectIntoTextFileContext) any

	// Visit a parse tree produced by MySqlParser#selectFieldsInto.
	VisitSelectFieldsInto(ctx *SelectFieldsIntoContext) any

	// Visit a parse tree produced by MySqlParser#selectLinesInto.
	VisitSelectLinesInto(ctx *SelectLinesIntoContext) any

	// Visit a parse tree produced by MySqlParser#fromClause.
	VisitFromClause(ctx *FromClauseContext) any

	// Visit a parse tree produced by MySqlParser#groupByClause.
	VisitGroupByClause(ctx *GroupByClauseContext) any

	// Visit a parse tree produced by MySqlParser#havingClause.
	VisitHavingClause(ctx *HavingClauseContext) any

	// Visit a parse tree produced by MySqlParser#groupByItem.
	VisitGroupByItem(ctx *GroupByItemContext) any

	// Visit a parse tree produced by MySqlParser#limitClause.
	VisitLimitClause(ctx *LimitClauseContext) any

	// Visit a parse tree produced by MySqlParser#limitClauseAtom.
	VisitLimitClauseAtom(ctx *LimitClauseAtomContext) any

	// Visit a parse tree produced by MySqlParser#startTransaction.
	VisitStartTransaction(ctx *StartTransactionContext) any

	// Visit a parse tree produced by MySqlParser#beginWork.
	VisitBeginWork(ctx *BeginWorkContext) any

	// Visit a parse tree produced by MySqlParser#commitWork.
	VisitCommitWork(ctx *CommitWorkContext) any

	// Visit a parse tree produced by MySqlParser#rollbackWork.
	VisitRollbackWork(ctx *RollbackWorkContext) any

	// Visit a parse tree produced by MySqlParser#savepointStatement.
	VisitSavepointStatement(ctx *SavepointStatementContext) any

	// Visit a parse tree produced by MySqlParser#rollbackStatement.
	VisitRollbackStatement(ctx *RollbackStatementContext) any

	// Visit a parse tree produced by MySqlParser#releaseStatement.
	VisitReleaseStatement(ctx *ReleaseStatementContext) any

	// Visit a parse tree produced by MySqlParser#lockTables.
	VisitLockTables(ctx *LockTablesContext) any

	// Visit a parse tree produced by MySqlParser#unlockTables.
	VisitUnlockTables(ctx *UnlockTablesContext) any

	// Visit a parse tree produced by MySqlParser#setAutocommitStatement.
	VisitSetAutocommitStatement(ctx *SetAutocommitStatementContext) any

	// Visit a parse tree produced by MySqlParser#setTransactionStatement.
	VisitSetTransactionStatement(ctx *SetTransactionStatementContext) any

	// Visit a parse tree produced by MySqlParser#transactionMode.
	VisitTransactionMode(ctx *TransactionModeContext) any

	// Visit a parse tree produced by MySqlParser#lockTableElement.
	VisitLockTableElement(ctx *LockTableElementContext) any

	// Visit a parse tree produced by MySqlParser#lockAction.
	VisitLockAction(ctx *LockActionContext) any

	// Visit a parse tree produced by MySqlParser#transactionOption.
	VisitTransactionOption(ctx *TransactionOptionContext) any

	// Visit a parse tree produced by MySqlParser#transactionLevel.
	VisitTransactionLevel(ctx *TransactionLevelContext) any

	// Visit a parse tree produced by MySqlParser#changeMaster.
	VisitChangeMaster(ctx *ChangeMasterContext) any

	// Visit a parse tree produced by MySqlParser#changeReplicationFilter.
	VisitChangeReplicationFilter(ctx *ChangeReplicationFilterContext) any

	// Visit a parse tree produced by MySqlParser#purgeBinaryLogs.
	VisitPurgeBinaryLogs(ctx *PurgeBinaryLogsContext) any

	// Visit a parse tree produced by MySqlParser#resetMaster.
	VisitResetMaster(ctx *ResetMasterContext) any

	// Visit a parse tree produced by MySqlParser#resetSlave.
	VisitResetSlave(ctx *ResetSlaveContext) any

	// Visit a parse tree produced by MySqlParser#startSlave.
	VisitStartSlave(ctx *StartSlaveContext) any

	// Visit a parse tree produced by MySqlParser#stopSlave.
	VisitStopSlave(ctx *StopSlaveContext) any

	// Visit a parse tree produced by MySqlParser#startGroupReplication.
	VisitStartGroupReplication(ctx *StartGroupReplicationContext) any

	// Visit a parse tree produced by MySqlParser#stopGroupReplication.
	VisitStopGroupReplication(ctx *StopGroupReplicationContext) any

	// Visit a parse tree produced by MySqlParser#masterStringOption.
	VisitMasterStringOption(ctx *MasterStringOptionContext) any

	// Visit a parse tree produced by MySqlParser#masterDecimalOption.
	VisitMasterDecimalOption(ctx *MasterDecimalOptionContext) any

	// Visit a parse tree produced by MySqlParser#masterBoolOption.
	VisitMasterBoolOption(ctx *MasterBoolOptionContext) any

	// Visit a parse tree produced by MySqlParser#masterRealOption.
	VisitMasterRealOption(ctx *MasterRealOptionContext) any

	// Visit a parse tree produced by MySqlParser#masterUidListOption.
	VisitMasterUidListOption(ctx *MasterUidListOptionContext) any

	// Visit a parse tree produced by MySqlParser#stringMasterOption.
	VisitStringMasterOption(ctx *StringMasterOptionContext) any

	// Visit a parse tree produced by MySqlParser#decimalMasterOption.
	VisitDecimalMasterOption(ctx *DecimalMasterOptionContext) any

	// Visit a parse tree produced by MySqlParser#boolMasterOption.
	VisitBoolMasterOption(ctx *BoolMasterOptionContext) any

	// Visit a parse tree produced by MySqlParser#channelOption.
	VisitChannelOption(ctx *ChannelOptionContext) any

	// Visit a parse tree produced by MySqlParser#doDbReplication.
	VisitDoDbReplication(ctx *DoDbReplicationContext) any

	// Visit a parse tree produced by MySqlParser#ignoreDbReplication.
	VisitIgnoreDbReplication(ctx *IgnoreDbReplicationContext) any

	// Visit a parse tree produced by MySqlParser#doTableReplication.
	VisitDoTableReplication(ctx *DoTableReplicationContext) any

	// Visit a parse tree produced by MySqlParser#ignoreTableReplication.
	VisitIgnoreTableReplication(ctx *IgnoreTableReplicationContext) any

	// Visit a parse tree produced by MySqlParser#wildDoTableReplication.
	VisitWildDoTableReplication(ctx *WildDoTableReplicationContext) any

	// Visit a parse tree produced by MySqlParser#wildIgnoreTableReplication.
	VisitWildIgnoreTableReplication(ctx *WildIgnoreTableReplicationContext) any

	// Visit a parse tree produced by MySqlParser#rewriteDbReplication.
	VisitRewriteDbReplication(ctx *RewriteDbReplicationContext) any

	// Visit a parse tree produced by MySqlParser#tablePair.
	VisitTablePair(ctx *TablePairContext) any

	// Visit a parse tree produced by MySqlParser#threadType.
	VisitThreadType(ctx *ThreadTypeContext) any

	// Visit a parse tree produced by MySqlParser#gtidsUntilOption.
	VisitGtidsUntilOption(ctx *GtidsUntilOptionContext) any

	// Visit a parse tree produced by MySqlParser#masterLogUntilOption.
	VisitMasterLogUntilOption(ctx *MasterLogUntilOptionContext) any

	// Visit a parse tree produced by MySqlParser#relayLogUntilOption.
	VisitRelayLogUntilOption(ctx *RelayLogUntilOptionContext) any

	// Visit a parse tree produced by MySqlParser#sqlGapsUntilOption.
	VisitSqlGapsUntilOption(ctx *SqlGapsUntilOptionContext) any

	// Visit a parse tree produced by MySqlParser#userConnectionOption.
	VisitUserConnectionOption(ctx *UserConnectionOptionContext) any

	// Visit a parse tree produced by MySqlParser#passwordConnectionOption.
	VisitPasswordConnectionOption(ctx *PasswordConnectionOptionContext) any

	// Visit a parse tree produced by MySqlParser#defaultAuthConnectionOption.
	VisitDefaultAuthConnectionOption(ctx *DefaultAuthConnectionOptionContext) any

	// Visit a parse tree produced by MySqlParser#pluginDirConnectionOption.
	VisitPluginDirConnectionOption(ctx *PluginDirConnectionOptionContext) any

	// Visit a parse tree produced by MySqlParser#gtuidSet.
	VisitGtuidSet(ctx *GtuidSetContext) any

	// Visit a parse tree produced by MySqlParser#xaStartTransaction.
	VisitXaStartTransaction(ctx *XaStartTransactionContext) any

	// Visit a parse tree produced by MySqlParser#xaEndTransaction.
	VisitXaEndTransaction(ctx *XaEndTransactionContext) any

	// Visit a parse tree produced by MySqlParser#xaPrepareStatement.
	VisitXaPrepareStatement(ctx *XaPrepareStatementContext) any

	// Visit a parse tree produced by MySqlParser#xaCommitWork.
	VisitXaCommitWork(ctx *XaCommitWorkContext) any

	// Visit a parse tree produced by MySqlParser#xaRollbackWork.
	VisitXaRollbackWork(ctx *XaRollbackWorkContext) any

	// Visit a parse tree produced by MySqlParser#xaRecoverWork.
	VisitXaRecoverWork(ctx *XaRecoverWorkContext) any

	// Visit a parse tree produced by MySqlParser#prepareStatement.
	VisitPrepareStatement(ctx *PrepareStatementContext) any

	// Visit a parse tree produced by MySqlParser#executeStatement.
	VisitExecuteStatement(ctx *ExecuteStatementContext) any

	// Visit a parse tree produced by MySqlParser#deallocatePrepare.
	VisitDeallocatePrepare(ctx *DeallocatePrepareContext) any

	// Visit a parse tree produced by MySqlParser#routineBody.
	VisitRoutineBody(ctx *RoutineBodyContext) any

	// Visit a parse tree produced by MySqlParser#blockStatement.
	VisitBlockStatement(ctx *BlockStatementContext) any

	// Visit a parse tree produced by MySqlParser#caseStatement.
	VisitCaseStatement(ctx *CaseStatementContext) any

	// Visit a parse tree produced by MySqlParser#ifStatement.
	VisitIfStatement(ctx *IfStatementContext) any

	// Visit a parse tree produced by MySqlParser#iterateStatement.
	VisitIterateStatement(ctx *IterateStatementContext) any

	// Visit a parse tree produced by MySqlParser#leaveStatement.
	VisitLeaveStatement(ctx *LeaveStatementContext) any

	// Visit a parse tree produced by MySqlParser#loopStatement.
	VisitLoopStatement(ctx *LoopStatementContext) any

	// Visit a parse tree produced by MySqlParser#repeatStatement.
	VisitRepeatStatement(ctx *RepeatStatementContext) any

	// Visit a parse tree produced by MySqlParser#returnStatement.
	VisitReturnStatement(ctx *ReturnStatementContext) any

	// Visit a parse tree produced by MySqlParser#whileStatement.
	VisitWhileStatement(ctx *WhileStatementContext) any

	// Visit a parse tree produced by MySqlParser#CloseCursor.
	VisitCloseCursor(ctx *CloseCursorContext) any

	// Visit a parse tree produced by MySqlParser#FetchCursor.
	VisitFetchCursor(ctx *FetchCursorContext) any

	// Visit a parse tree produced by MySqlParser#OpenCursor.
	VisitOpenCursor(ctx *OpenCursorContext) any

	// Visit a parse tree produced by MySqlParser#declareVariable.
	VisitDeclareVariable(ctx *DeclareVariableContext) any

	// Visit a parse tree produced by MySqlParser#declareCondition.
	VisitDeclareCondition(ctx *DeclareConditionContext) any

	// Visit a parse tree produced by MySqlParser#declareCursor.
	VisitDeclareCursor(ctx *DeclareCursorContext) any

	// Visit a parse tree produced by MySqlParser#declareHandler.
	VisitDeclareHandler(ctx *DeclareHandlerContext) any

	// Visit a parse tree produced by MySqlParser#handlerConditionCode.
	VisitHandlerConditionCode(ctx *HandlerConditionCodeContext) any

	// Visit a parse tree produced by MySqlParser#handlerConditionState.
	VisitHandlerConditionState(ctx *HandlerConditionStateContext) any

	// Visit a parse tree produced by MySqlParser#handlerConditionName.
	VisitHandlerConditionName(ctx *HandlerConditionNameContext) any

	// Visit a parse tree produced by MySqlParser#handlerConditionWarning.
	VisitHandlerConditionWarning(ctx *HandlerConditionWarningContext) any

	// Visit a parse tree produced by MySqlParser#handlerConditionNotfound.
	VisitHandlerConditionNotfound(ctx *HandlerConditionNotfoundContext) any

	// Visit a parse tree produced by MySqlParser#handlerConditionException.
	VisitHandlerConditionException(ctx *HandlerConditionExceptionContext) any

	// Visit a parse tree produced by MySqlParser#procedureSqlStatement.
	VisitProcedureSqlStatement(ctx *ProcedureSqlStatementContext) any

	// Visit a parse tree produced by MySqlParser#caseAlternative.
	VisitCaseAlternative(ctx *CaseAlternativeContext) any

	// Visit a parse tree produced by MySqlParser#elifAlternative.
	VisitElifAlternative(ctx *ElifAlternativeContext) any

	// Visit a parse tree produced by MySqlParser#alterUserMysqlV56.
	VisitAlterUserMysqlV56(ctx *AlterUserMysqlV56Context) any

	// Visit a parse tree produced by MySqlParser#alterUserMysqlV57.
	VisitAlterUserMysqlV57(ctx *AlterUserMysqlV57Context) any

	// Visit a parse tree produced by MySqlParser#createUserMysqlV56.
	VisitCreateUserMysqlV56(ctx *CreateUserMysqlV56Context) any

	// Visit a parse tree produced by MySqlParser#createUserMysqlV57.
	VisitCreateUserMysqlV57(ctx *CreateUserMysqlV57Context) any

	// Visit a parse tree produced by MySqlParser#dropUser.
	VisitDropUser(ctx *DropUserContext) any

	// Visit a parse tree produced by MySqlParser#grantStatement.
	VisitGrantStatement(ctx *GrantStatementContext) any

	// Visit a parse tree produced by MySqlParser#grantProxy.
	VisitGrantProxy(ctx *GrantProxyContext) any

	// Visit a parse tree produced by MySqlParser#renameUser.
	VisitRenameUser(ctx *RenameUserContext) any

	// Visit a parse tree produced by MySqlParser#detailRevoke.
	VisitDetailRevoke(ctx *DetailRevokeContext) any

	// Visit a parse tree produced by MySqlParser#shortRevoke.
	VisitShortRevoke(ctx *ShortRevokeContext) any

	// Visit a parse tree produced by MySqlParser#revokeProxy.
	VisitRevokeProxy(ctx *RevokeProxyContext) any

	// Visit a parse tree produced by MySqlParser#setPasswordStatement.
	VisitSetPasswordStatement(ctx *SetPasswordStatementContext) any

	// Visit a parse tree produced by MySqlParser#userSpecification.
	VisitUserSpecification(ctx *UserSpecificationContext) any

	// Visit a parse tree produced by MySqlParser#passwordAuthOption.
	VisitPasswordAuthOption(ctx *PasswordAuthOptionContext) any

	// Visit a parse tree produced by MySqlParser#stringAuthOption.
	VisitStringAuthOption(ctx *StringAuthOptionContext) any

	// Visit a parse tree produced by MySqlParser#hashAuthOption.
	VisitHashAuthOption(ctx *HashAuthOptionContext) any

	// Visit a parse tree produced by MySqlParser#simpleAuthOption.
	VisitSimpleAuthOption(ctx *SimpleAuthOptionContext) any

	// Visit a parse tree produced by MySqlParser#tlsOption.
	VisitTlsOption(ctx *TlsOptionContext) any

	// Visit a parse tree produced by MySqlParser#userResourceOption.
	VisitUserResourceOption(ctx *UserResourceOptionContext) any

	// Visit a parse tree produced by MySqlParser#userPasswordOption.
	VisitUserPasswordOption(ctx *UserPasswordOptionContext) any

	// Visit a parse tree produced by MySqlParser#userLockOption.
	VisitUserLockOption(ctx *UserLockOptionContext) any

	// Visit a parse tree produced by MySqlParser#privelegeClause.
	VisitPrivelegeClause(ctx *PrivelegeClauseContext) any

	// Visit a parse tree produced by MySqlParser#privilege.
	VisitPrivilege(ctx *PrivilegeContext) any

	// Visit a parse tree produced by MySqlParser#currentSchemaPriviLevel.
	VisitCurrentSchemaPriviLevel(ctx *CurrentSchemaPriviLevelContext) any

	// Visit a parse tree produced by MySqlParser#globalPrivLevel.
	VisitGlobalPrivLevel(ctx *GlobalPrivLevelContext) any

	// Visit a parse tree produced by MySqlParser#definiteSchemaPrivLevel.
	VisitDefiniteSchemaPrivLevel(ctx *DefiniteSchemaPrivLevelContext) any

	// Visit a parse tree produced by MySqlParser#definiteFullTablePrivLevel.
	VisitDefiniteFullTablePrivLevel(ctx *DefiniteFullTablePrivLevelContext) any

	// Visit a parse tree produced by MySqlParser#definiteFullTablePrivLevel2.
	VisitDefiniteFullTablePrivLevel2(ctx *DefiniteFullTablePrivLevel2Context) any

	// Visit a parse tree produced by MySqlParser#definiteTablePrivLevel.
	VisitDefiniteTablePrivLevel(ctx *DefiniteTablePrivLevelContext) any

	// Visit a parse tree produced by MySqlParser#renameUserClause.
	VisitRenameUserClause(ctx *RenameUserClauseContext) any

	// Visit a parse tree produced by MySqlParser#analyzeTable.
	VisitAnalyzeTable(ctx *AnalyzeTableContext) any

	// Visit a parse tree produced by MySqlParser#checkTable.
	VisitCheckTable(ctx *CheckTableContext) any

	// Visit a parse tree produced by MySqlParser#checksumTable.
	VisitChecksumTable(ctx *ChecksumTableContext) any

	// Visit a parse tree produced by MySqlParser#optimizeTable.
	VisitOptimizeTable(ctx *OptimizeTableContext) any

	// Visit a parse tree produced by MySqlParser#repairTable.
	VisitRepairTable(ctx *RepairTableContext) any

	// Visit a parse tree produced by MySqlParser#checkTableOption.
	VisitCheckTableOption(ctx *CheckTableOptionContext) any

	// Visit a parse tree produced by MySqlParser#createUdfunction.
	VisitCreateUdfunction(ctx *CreateUdfunctionContext) any

	// Visit a parse tree produced by MySqlParser#installPlugin.
	VisitInstallPlugin(ctx *InstallPluginContext) any

	// Visit a parse tree produced by MySqlParser#uninstallPlugin.
	VisitUninstallPlugin(ctx *UninstallPluginContext) any

	// Visit a parse tree produced by MySqlParser#setVariable.
	VisitSetVariable(ctx *SetVariableContext) any

	// Visit a parse tree produced by MySqlParser#setCharset.
	VisitSetCharset(ctx *SetCharsetContext) any

	// Visit a parse tree produced by MySqlParser#setNames.
	VisitSetNames(ctx *SetNamesContext) any

	// Visit a parse tree produced by MySqlParser#setPassword.
	VisitSetPassword(ctx *SetPasswordContext) any

	// Visit a parse tree produced by MySqlParser#setTransaction.
	VisitSetTransaction(ctx *SetTransactionContext) any

	// Visit a parse tree produced by MySqlParser#setAutocommit.
	VisitSetAutocommit(ctx *SetAutocommitContext) any

	// Visit a parse tree produced by MySqlParser#setNewValueInsideTrigger.
	VisitSetNewValueInsideTrigger(ctx *SetNewValueInsideTriggerContext) any

	// Visit a parse tree produced by MySqlParser#showMasterLogs.
	VisitShowMasterLogs(ctx *ShowMasterLogsContext) any

	// Visit a parse tree produced by MySqlParser#showLogEvents.
	VisitShowLogEvents(ctx *ShowLogEventsContext) any

	// Visit a parse tree produced by MySqlParser#showObjectFilter.
	VisitShowObjectFilter(ctx *ShowObjectFilterContext) any

	// Visit a parse tree produced by MySqlParser#showColumns.
	VisitShowColumns(ctx *ShowColumnsContext) any

	// Visit a parse tree produced by MySqlParser#showCreateDb.
	VisitShowCreateDb(ctx *ShowCreateDbContext) any

	// Visit a parse tree produced by MySqlParser#showCreateFullIdObject.
	VisitShowCreateFullIdObject(ctx *ShowCreateFullIdObjectContext) any

	// Visit a parse tree produced by MySqlParser#showCreateUser.
	VisitShowCreateUser(ctx *ShowCreateUserContext) any

	// Visit a parse tree produced by MySqlParser#showEngine.
	VisitShowEngine(ctx *ShowEngineContext) any

	// Visit a parse tree produced by MySqlParser#showGlobalInfo.
	VisitShowGlobalInfo(ctx *ShowGlobalInfoContext) any

	// Visit a parse tree produced by MySqlParser#showErrors.
	VisitShowErrors(ctx *ShowErrorsContext) any

	// Visit a parse tree produced by MySqlParser#showCountErrors.
	VisitShowCountErrors(ctx *ShowCountErrorsContext) any

	// Visit a parse tree produced by MySqlParser#showSchemaFilter.
	VisitShowSchemaFilter(ctx *ShowSchemaFilterContext) any

	// Visit a parse tree produced by MySqlParser#showRoutine.
	VisitShowRoutine(ctx *ShowRoutineContext) any

	// Visit a parse tree produced by MySqlParser#showGrants.
	VisitShowGrants(ctx *ShowGrantsContext) any

	// Visit a parse tree produced by MySqlParser#showIndexes.
	VisitShowIndexes(ctx *ShowIndexesContext) any

	// Visit a parse tree produced by MySqlParser#showOpenTables.
	VisitShowOpenTables(ctx *ShowOpenTablesContext) any

	// Visit a parse tree produced by MySqlParser#showProfile.
	VisitShowProfile(ctx *ShowProfileContext) any

	// Visit a parse tree produced by MySqlParser#showSlaveStatus.
	VisitShowSlaveStatus(ctx *ShowSlaveStatusContext) any

	// Visit a parse tree produced by MySqlParser#variableClause.
	VisitVariableClause(ctx *VariableClauseContext) any

	// Visit a parse tree produced by MySqlParser#showCommonEntity.
	VisitShowCommonEntity(ctx *ShowCommonEntityContext) any

	// Visit a parse tree produced by MySqlParser#showFilter.
	VisitShowFilter(ctx *ShowFilterContext) any

	// Visit a parse tree produced by MySqlParser#showGlobalInfoClause.
	VisitShowGlobalInfoClause(ctx *ShowGlobalInfoClauseContext) any

	// Visit a parse tree produced by MySqlParser#showSchemaEntity.
	VisitShowSchemaEntity(ctx *ShowSchemaEntityContext) any

	// Visit a parse tree produced by MySqlParser#showProfileType.
	VisitShowProfileType(ctx *ShowProfileTypeContext) any

	// Visit a parse tree produced by MySqlParser#binlogStatement.
	VisitBinlogStatement(ctx *BinlogStatementContext) any

	// Visit a parse tree produced by MySqlParser#cacheIndexStatement.
	VisitCacheIndexStatement(ctx *CacheIndexStatementContext) any

	// Visit a parse tree produced by MySqlParser#flushStatement.
	VisitFlushStatement(ctx *FlushStatementContext) any

	// Visit a parse tree produced by MySqlParser#killStatement.
	VisitKillStatement(ctx *KillStatementContext) any

	// Visit a parse tree produced by MySqlParser#loadIndexIntoCache.
	VisitLoadIndexIntoCache(ctx *LoadIndexIntoCacheContext) any

	// Visit a parse tree produced by MySqlParser#resetStatement.
	VisitResetStatement(ctx *ResetStatementContext) any

	// Visit a parse tree produced by MySqlParser#shutdownStatement.
	VisitShutdownStatement(ctx *ShutdownStatementContext) any

	// Visit a parse tree produced by MySqlParser#tableIndexes.
	VisitTableIndexes(ctx *TableIndexesContext) any

	// Visit a parse tree produced by MySqlParser#simpleFlushOption.
	VisitSimpleFlushOption(ctx *SimpleFlushOptionContext) any

	// Visit a parse tree produced by MySqlParser#channelFlushOption.
	VisitChannelFlushOption(ctx *ChannelFlushOptionContext) any

	// Visit a parse tree produced by MySqlParser#tableFlushOption.
	VisitTableFlushOption(ctx *TableFlushOptionContext) any

	// Visit a parse tree produced by MySqlParser#flushTableOption.
	VisitFlushTableOption(ctx *FlushTableOptionContext) any

	// Visit a parse tree produced by MySqlParser#loadedTableIndexes.
	VisitLoadedTableIndexes(ctx *LoadedTableIndexesContext) any

	// Visit a parse tree produced by MySqlParser#simpleDescribeStatement.
	VisitSimpleDescribeStatement(ctx *SimpleDescribeStatementContext) any

	// Visit a parse tree produced by MySqlParser#fullDescribeStatement.
	VisitFullDescribeStatement(ctx *FullDescribeStatementContext) any

	// Visit a parse tree produced by MySqlParser#helpStatement.
	VisitHelpStatement(ctx *HelpStatementContext) any

	// Visit a parse tree produced by MySqlParser#useStatement.
	VisitUseStatement(ctx *UseStatementContext) any

	// Visit a parse tree produced by MySqlParser#signalStatement.
	VisitSignalStatement(ctx *SignalStatementContext) any

	// Visit a parse tree produced by MySqlParser#resignalStatement.
	VisitResignalStatement(ctx *ResignalStatementContext) any

	// Visit a parse tree produced by MySqlParser#signalConditionInformation.
	VisitSignalConditionInformation(ctx *SignalConditionInformationContext) any

	// Visit a parse tree produced by MySqlParser#diagnosticsStatement.
	VisitDiagnosticsStatement(ctx *DiagnosticsStatementContext) any

	// Visit a parse tree produced by MySqlParser#diagnosticsConditionInformationName.
	VisitDiagnosticsConditionInformationName(ctx *DiagnosticsConditionInformationNameContext) any

	// Visit a parse tree produced by MySqlParser#describeStatements.
	VisitDescribeStatements(ctx *DescribeStatementsContext) any

	// Visit a parse tree produced by MySqlParser#describeConnection.
	VisitDescribeConnection(ctx *DescribeConnectionContext) any

	// Visit a parse tree produced by MySqlParser#fullId.
	VisitFullId(ctx *FullIdContext) any

	// Visit a parse tree produced by MySqlParser#tableName.
	VisitTableName(ctx *TableNameContext) any

	// Visit a parse tree produced by MySqlParser#fullColumnName.
	VisitFullColumnName(ctx *FullColumnNameContext) any

	// Visit a parse tree produced by MySqlParser#indexColumnName.
	VisitIndexColumnName(ctx *IndexColumnNameContext) any

	// Visit a parse tree produced by MySqlParser#userName.
	VisitUserName(ctx *UserNameContext) any

	// Visit a parse tree produced by MySqlParser#mysqlVariable.
	VisitMysqlVariable(ctx *MysqlVariableContext) any

	// Visit a parse tree produced by MySqlParser#charsetName.
	VisitCharsetName(ctx *CharsetNameContext) any

	// Visit a parse tree produced by MySqlParser#collationName.
	VisitCollationName(ctx *CollationNameContext) any

	// Visit a parse tree produced by MySqlParser#engineName.
	VisitEngineName(ctx *EngineNameContext) any

	// Visit a parse tree produced by MySqlParser#uuidSet.
	VisitUuidSet(ctx *UuidSetContext) any

	// Visit a parse tree produced by MySqlParser#xid.
	VisitXid(ctx *XidContext) any

	// Visit a parse tree produced by MySqlParser#xuidStringId.
	VisitXuidStringId(ctx *XuidStringIdContext) any

	// Visit a parse tree produced by MySqlParser#authPlugin.
	VisitAuthPlugin(ctx *AuthPluginContext) any

	// Visit a parse tree produced by MySqlParser#uid.
	VisitUid(ctx *UidContext) any

	// Visit a parse tree produced by MySqlParser#simpleId.
	VisitSimpleId(ctx *SimpleIdContext) any

	// Visit a parse tree produced by MySqlParser#dottedId.
	VisitDottedId(ctx *DottedIdContext) any

	// Visit a parse tree produced by MySqlParser#decimalLiteral.
	VisitDecimalLiteral(ctx *DecimalLiteralContext) any

	// Visit a parse tree produced by MySqlParser#fileSizeLiteral.
	VisitFileSizeLiteral(ctx *FileSizeLiteralContext) any

	// Visit a parse tree produced by MySqlParser#stringLiteral.
	VisitStringLiteral(ctx *StringLiteralContext) any

	// Visit a parse tree produced by MySqlParser#booleanLiteral.
	VisitBooleanLiteral(ctx *BooleanLiteralContext) any

	// Visit a parse tree produced by MySqlParser#hexadecimalLiteral.
	VisitHexadecimalLiteral(ctx *HexadecimalLiteralContext) any

	// Visit a parse tree produced by MySqlParser#nullNotnull.
	VisitNullNotnull(ctx *NullNotnullContext) any

	// Visit a parse tree produced by MySqlParser#constant.
	VisitConstant(ctx *ConstantContext) any

	// Visit a parse tree produced by MySqlParser#stringDataType.
	VisitStringDataType(ctx *StringDataTypeContext) any

	// Visit a parse tree produced by MySqlParser#nationalStringDataType.
	VisitNationalStringDataType(ctx *NationalStringDataTypeContext) any

	// Visit a parse tree produced by MySqlParser#nationalVaryingStringDataType.
	VisitNationalVaryingStringDataType(ctx *NationalVaryingStringDataTypeContext) any

	// Visit a parse tree produced by MySqlParser#dimensionDataType.
	VisitDimensionDataType(ctx *DimensionDataTypeContext) any

	// Visit a parse tree produced by MySqlParser#simpleDataType.
	VisitSimpleDataType(ctx *SimpleDataTypeContext) any

	// Visit a parse tree produced by MySqlParser#collectionDataType.
	VisitCollectionDataType(ctx *CollectionDataTypeContext) any

	// Visit a parse tree produced by MySqlParser#spatialDataType.
	VisitSpatialDataType(ctx *SpatialDataTypeContext) any

	// Visit a parse tree produced by MySqlParser#longVarcharDataType.
	VisitLongVarcharDataType(ctx *LongVarcharDataTypeContext) any

	// Visit a parse tree produced by MySqlParser#longVarbinaryDataType.
	VisitLongVarbinaryDataType(ctx *LongVarbinaryDataTypeContext) any

	// Visit a parse tree produced by MySqlParser#collectionOptions.
	VisitCollectionOptions(ctx *CollectionOptionsContext) any

	// Visit a parse tree produced by MySqlParser#convertedDataType.
	VisitConvertedDataType(ctx *ConvertedDataTypeContext) any

	// Visit a parse tree produced by MySqlParser#lengthOneDimension.
	VisitLengthOneDimension(ctx *LengthOneDimensionContext) any

	// Visit a parse tree produced by MySqlParser#lengthTwoDimension.
	VisitLengthTwoDimension(ctx *LengthTwoDimensionContext) any

	// Visit a parse tree produced by MySqlParser#lengthTwoOptionalDimension.
	VisitLengthTwoOptionalDimension(ctx *LengthTwoOptionalDimensionContext) any

	// Visit a parse tree produced by MySqlParser#uidList.
	VisitUidList(ctx *UidListContext) any

	// Visit a parse tree produced by MySqlParser#tables.
	VisitTables(ctx *TablesContext) any

	// Visit a parse tree produced by MySqlParser#indexColumnNames.
	VisitIndexColumnNames(ctx *IndexColumnNamesContext) any

	// Visit a parse tree produced by MySqlParser#expressions.
	VisitExpressions(ctx *ExpressionsContext) any

	// Visit a parse tree produced by MySqlParser#expressionsWithDefaults.
	VisitExpressionsWithDefaults(ctx *ExpressionsWithDefaultsContext) any

	// Visit a parse tree produced by MySqlParser#constants.
	VisitConstants(ctx *ConstantsContext) any

	// Visit a parse tree produced by MySqlParser#simpleStrings.
	VisitSimpleStrings(ctx *SimpleStringsContext) any

	// Visit a parse tree produced by MySqlParser#userVariables.
	VisitUserVariables(ctx *UserVariablesContext) any

	// Visit a parse tree produced by MySqlParser#defaultValue.
	VisitDefaultValue(ctx *DefaultValueContext) any

	// Visit a parse tree produced by MySqlParser#currentTimestamp.
	VisitCurrentTimestamp(ctx *CurrentTimestampContext) any

	// Visit a parse tree produced by MySqlParser#expressionOrDefault.
	VisitExpressionOrDefault(ctx *ExpressionOrDefaultContext) any

	// Visit a parse tree produced by MySqlParser#ifExists.
	VisitIfExists(ctx *IfExistsContext) any

	// Visit a parse tree produced by MySqlParser#ifNotExists.
	VisitIfNotExists(ctx *IfNotExistsContext) any

	// Visit a parse tree produced by MySqlParser#specificFunctionCall.
	VisitSpecificFunctionCall(ctx *SpecificFunctionCallContext) any

	// Visit a parse tree produced by MySqlParser#aggregateFunctionCall.
	VisitAggregateFunctionCall(ctx *AggregateFunctionCallContext) any

	// Visit a parse tree produced by MySqlParser#scalarFunctionCall.
	VisitScalarFunctionCall(ctx *ScalarFunctionCallContext) any

	// Visit a parse tree produced by MySqlParser#udfFunctionCall.
	VisitUdfFunctionCall(ctx *UdfFunctionCallContext) any

	// Visit a parse tree produced by MySqlParser#passwordFunctionCall.
	VisitPasswordFunctionCall(ctx *PasswordFunctionCallContext) any

	// Visit a parse tree produced by MySqlParser#simpleFunctionCall.
	VisitSimpleFunctionCall(ctx *SimpleFunctionCallContext) any

	// Visit a parse tree produced by MySqlParser#dataTypeFunctionCall.
	VisitDataTypeFunctionCall(ctx *DataTypeFunctionCallContext) any

	// Visit a parse tree produced by MySqlParser#valuesFunctionCall.
	VisitValuesFunctionCall(ctx *ValuesFunctionCallContext) any

	// Visit a parse tree produced by MySqlParser#caseExpressionFunctionCall.
	VisitCaseExpressionFunctionCall(ctx *CaseExpressionFunctionCallContext) any

	// Visit a parse tree produced by MySqlParser#caseFunctionCall.
	VisitCaseFunctionCall(ctx *CaseFunctionCallContext) any

	// Visit a parse tree produced by MySqlParser#charFunctionCall.
	VisitCharFunctionCall(ctx *CharFunctionCallContext) any

	// Visit a parse tree produced by MySqlParser#positionFunctionCall.
	VisitPositionFunctionCall(ctx *PositionFunctionCallContext) any

	// Visit a parse tree produced by MySqlParser#substrFunctionCall.
	VisitSubstrFunctionCall(ctx *SubstrFunctionCallContext) any

	// Visit a parse tree produced by MySqlParser#trimFunctionCall.
	VisitTrimFunctionCall(ctx *TrimFunctionCallContext) any

	// Visit a parse tree produced by MySqlParser#weightFunctionCall.
	VisitWeightFunctionCall(ctx *WeightFunctionCallContext) any

	// Visit a parse tree produced by MySqlParser#extractFunctionCall.
	VisitExtractFunctionCall(ctx *ExtractFunctionCallContext) any

	// Visit a parse tree produced by MySqlParser#getFormatFunctionCall.
	VisitGetFormatFunctionCall(ctx *GetFormatFunctionCallContext) any

	// Visit a parse tree produced by MySqlParser#jsonValueFunctionCall.
	VisitJsonValueFunctionCall(ctx *JsonValueFunctionCallContext) any

	// Visit a parse tree produced by MySqlParser#caseFuncAlternative.
	VisitCaseFuncAlternative(ctx *CaseFuncAlternativeContext) any

	// Visit a parse tree produced by MySqlParser#levelWeightList.
	VisitLevelWeightList(ctx *LevelWeightListContext) any

	// Visit a parse tree produced by MySqlParser#levelWeightRange.
	VisitLevelWeightRange(ctx *LevelWeightRangeContext) any

	// Visit a parse tree produced by MySqlParser#levelInWeightListElement.
	VisitLevelInWeightListElement(ctx *LevelInWeightListElementContext) any

	// Visit a parse tree produced by MySqlParser#aggregateWindowedFunction.
	VisitAggregateWindowedFunction(ctx *AggregateWindowedFunctionContext) any

	// Visit a parse tree produced by MySqlParser#scalarFunctionName.
	VisitScalarFunctionName(ctx *ScalarFunctionNameContext) any

	// Visit a parse tree produced by MySqlParser#passwordFunctionClause.
	VisitPasswordFunctionClause(ctx *PasswordFunctionClauseContext) any

	// Visit a parse tree produced by MySqlParser#functionArgs.
	VisitFunctionArgs(ctx *FunctionArgsContext) any

	// Visit a parse tree produced by MySqlParser#functionArg.
	VisitFunctionArg(ctx *FunctionArgContext) any

	// Visit a parse tree produced by MySqlParser#isExpression.
	VisitIsExpression(ctx *IsExpressionContext) any

	// Visit a parse tree produced by MySqlParser#notExpression.
	VisitNotExpression(ctx *NotExpressionContext) any

	// Visit a parse tree produced by MySqlParser#logicalExpression.
	VisitLogicalExpression(ctx *LogicalExpressionContext) any

	// Visit a parse tree produced by MySqlParser#predicateExpression.
	VisitPredicateExpression(ctx *PredicateExpressionContext) any

	// Visit a parse tree produced by MySqlParser#soundsLikePredicate.
	VisitSoundsLikePredicate(ctx *SoundsLikePredicateContext) any

	// Visit a parse tree produced by MySqlParser#expressionAtomPredicate.
	VisitExpressionAtomPredicate(ctx *ExpressionAtomPredicateContext) any

	// Visit a parse tree produced by MySqlParser#subqueryComparisonPredicate.
	VisitSubqueryComparisonPredicate(ctx *SubqueryComparisonPredicateContext) any

	// Visit a parse tree produced by MySqlParser#jsonMemberOfPredicate.
	VisitJsonMemberOfPredicate(ctx *JsonMemberOfPredicateContext) any

	// Visit a parse tree produced by MySqlParser#binaryComparisonPredicate.
	VisitBinaryComparisonPredicate(ctx *BinaryComparisonPredicateContext) any

	// Visit a parse tree produced by MySqlParser#inPredicate.
	VisitInPredicate(ctx *InPredicateContext) any

	// Visit a parse tree produced by MySqlParser#betweenPredicate.
	VisitBetweenPredicate(ctx *BetweenPredicateContext) any

	// Visit a parse tree produced by MySqlParser#isNullPredicate.
	VisitIsNullPredicate(ctx *IsNullPredicateContext) any

	// Visit a parse tree produced by MySqlParser#likePredicate.
	VisitLikePredicate(ctx *LikePredicateContext) any

	// Visit a parse tree produced by MySqlParser#regexpPredicate.
	VisitRegexpPredicate(ctx *RegexpPredicateContext) any

	// Visit a parse tree produced by MySqlParser#unaryExpressionAtom.
	VisitUnaryExpressionAtom(ctx *UnaryExpressionAtomContext) any

	// Visit a parse tree produced by MySqlParser#collateExpressionAtom.
	VisitCollateExpressionAtom(ctx *CollateExpressionAtomContext) any

	// Visit a parse tree produced by MySqlParser#mysqlVariableExpressionAtom.
	VisitMysqlVariableExpressionAtom(ctx *MysqlVariableExpressionAtomContext) any

	// Visit a parse tree produced by MySqlParser#nestedExpressionAtom.
	VisitNestedExpressionAtom(ctx *NestedExpressionAtomContext) any

	// Visit a parse tree produced by MySqlParser#nestedRowExpressionAtom.
	VisitNestedRowExpressionAtom(ctx *NestedRowExpressionAtomContext) any

	// Visit a parse tree produced by MySqlParser#mathExpressionAtom.
	VisitMathExpressionAtom(ctx *MathExpressionAtomContext) any

	// Visit a parse tree produced by MySqlParser#existsExpressionAtom.
	VisitExistsExpressionAtom(ctx *ExistsExpressionAtomContext) any

	// Visit a parse tree produced by MySqlParser#intervalExpressionAtom.
	VisitIntervalExpressionAtom(ctx *IntervalExpressionAtomContext) any

	// Visit a parse tree produced by MySqlParser#jsonExpressionAtom.
	VisitJsonExpressionAtom(ctx *JsonExpressionAtomContext) any

	// Visit a parse tree produced by MySqlParser#subqueryExpressionAtom.
	VisitSubqueryExpressionAtom(ctx *SubqueryExpressionAtomContext) any

	// Visit a parse tree produced by MySqlParser#constantExpressionAtom.
	VisitConstantExpressionAtom(ctx *ConstantExpressionAtomContext) any

	// Visit a parse tree produced by MySqlParser#functionCallExpressionAtom.
	VisitFunctionCallExpressionAtom(ctx *FunctionCallExpressionAtomContext) any

	// Visit a parse tree produced by MySqlParser#binaryExpressionAtom.
	VisitBinaryExpressionAtom(ctx *BinaryExpressionAtomContext) any

	// Visit a parse tree produced by MySqlParser#fullColumnNameExpressionAtom.
	VisitFullColumnNameExpressionAtom(ctx *FullColumnNameExpressionAtomContext) any

	// Visit a parse tree produced by MySqlParser#bitExpressionAtom.
	VisitBitExpressionAtom(ctx *BitExpressionAtomContext) any

	// Visit a parse tree produced by MySqlParser#unaryOperator.
	VisitUnaryOperator(ctx *UnaryOperatorContext) any

	// Visit a parse tree produced by MySqlParser#comparisonOperator.
	VisitComparisonOperator(ctx *ComparisonOperatorContext) any

	// Visit a parse tree produced by MySqlParser#logicalOperator.
	VisitLogicalOperator(ctx *LogicalOperatorContext) any

	// Visit a parse tree produced by MySqlParser#bitOperator.
	VisitBitOperator(ctx *BitOperatorContext) any

	// Visit a parse tree produced by MySqlParser#mathOperator.
	VisitMathOperator(ctx *MathOperatorContext) any

	// Visit a parse tree produced by MySqlParser#jsonOperator.
	VisitJsonOperator(ctx *JsonOperatorContext) any

	// Visit a parse tree produced by MySqlParser#charsetNameBase.
	VisitCharsetNameBase(ctx *CharsetNameBaseContext) any

	// Visit a parse tree produced by MySqlParser#transactionLevelBase.
	VisitTransactionLevelBase(ctx *TransactionLevelBaseContext) any

	// Visit a parse tree produced by MySqlParser#privilegesBase.
	VisitPrivilegesBase(ctx *PrivilegesBaseContext) any

	// Visit a parse tree produced by MySqlParser#intervalTypeBase.
	VisitIntervalTypeBase(ctx *IntervalTypeBaseContext) any

	// Visit a parse tree produced by MySqlParser#dataTypeBase.
	VisitDataTypeBase(ctx *DataTypeBaseContext) any

	// Visit a parse tree produced by MySqlParser#keywordsCanBeId.
	VisitKeywordsCanBeId(ctx *KeywordsCanBeIdContext) any

	// Visit a parse tree produced by MySqlParser#functionNameBase.
	VisitFunctionNameBase(ctx *FunctionNameBaseContext) any
}
