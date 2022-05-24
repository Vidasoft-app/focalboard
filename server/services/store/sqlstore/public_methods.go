// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

// Code generated by "make generate" from the Store interface
// DO NOT EDIT

// To add a public method, create an entry in the Store interface,
// prefix it with a @withTransaction comment if you need it to be
// transactional and then add a private method in the store itself
// with db sq.BaseRunner as the first parameter before running `make
// generate`

package sqlstore

import (
	"context"
	"time"

	mmModel "github.com/mattermost/mattermost-server/v6/model"

	"github.com/mattermost/focalboard/server/model"
	"github.com/mattermost/focalboard/server/services/store"
	"github.com/mattermost/mattermost-server/v6/shared/mlog"
)

func (s *SQLStore) CheckUserIDInTeam(userID string, teamID string) bool {
	return s.checkUserIDInTeam(s.db, userID, teamID)

}

func (s *SQLStore) CleanUpSessions(expireTime int64) error {
	return s.cleanUpSessions(s.db, expireTime)

}

func (s *SQLStore) CreatePrivateWorkspace(userID string) (string, error) {
	return s.createPrivateWorkspace(s.db, userID)

}

func (s *SQLStore) CreateSession(session *model.Session) error {
	return s.createSession(s.db, session)

}

func (s *SQLStore) CreateSubscription(c store.Container, sub *model.Subscription) (*model.Subscription, error) {
	return s.createSubscription(s.db, c, sub)

}

func (s *SQLStore) CreateUser(user *model.User) error {
	return s.createUser(s.db, user)

}

func (s *SQLStore) DeleteBlock(c store.Container, blockID string, modifiedBy string) error {
	if s.dbType == sqliteDBType {
		return s.deleteBlock(s.db, c, blockID, modifiedBy)
	}
	tx, txErr := s.db.BeginTx(context.Background(), nil)
	if txErr != nil {
		return txErr
	}
	err := s.deleteBlock(tx, c, blockID, modifiedBy)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			s.logger.Error("transaction rollback error", mlog.Err(rollbackErr), mlog.String("methodName", "DeleteBlock"))
		}
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil

}

func (s *SQLStore) DeleteNotificationHint(c store.Container, blockID string) error {
	return s.deleteNotificationHint(s.db, c, blockID)

}

func (s *SQLStore) DeleteSession(sessionID string) error {
	return s.deleteSession(s.db, sessionID)

}

func (s *SQLStore) DeleteSubscription(c store.Container, blockID string, subscriberID string) error {
	return s.deleteSubscription(s.db, c, blockID, subscriberID)

}

func (s *SQLStore) GetActiveUserCount(updatedSecondsAgo int64) (int, error) {
	return s.getActiveUserCount(s.db, updatedSecondsAgo)

}

func (s *SQLStore) GetAllBlocks(c store.Container) ([]model.Block, error) {
	return s.getAllBlocks(s.db, c)

}

func (s *SQLStore) GetBlock(c store.Container, blockID string) (*model.Block, error) {
	return s.getBlock(s.db, c, blockID)

}

func (s *SQLStore) GetBlockCountsByType() (map[string]int64, error) {
	return s.getBlockCountsByType(s.db)

}

func (s *SQLStore) GetBlockHistory(c store.Container, blockID string, opts model.QueryBlockHistoryOptions) ([]model.Block, error) {
	return s.getBlockHistory(s.db, c, blockID, opts)

}

func (s *SQLStore) GetBlockHistoryDescendants(boardID string, opts model.QueryBlockHistoryOptions) ([]model.Block, error) {
	return s.getBlockHistoryDescendants(s.db, boardID, opts)

}

func (s *SQLStore) GetBlocksWithParent(c store.Container, parentID string) ([]model.Block, error) {
	return s.getBlocksWithParent(s.db, c, parentID)

}

func (s *SQLStore) GetBlocksWithParentAndType(c store.Container, parentID string, blockType string) ([]model.Block, error) {
	return s.getBlocksWithParentAndType(s.db, c, parentID, blockType)

}

func (s *SQLStore) GetBlocksWithRootID(c store.Container, rootID string) ([]model.Block, error) {
	return s.getBlocksWithRootID(s.db, c, rootID)

}

func (s *SQLStore) GetBlocksWithType(c store.Container, blockType string) ([]model.Block, error) {
	return s.getBlocksWithType(s.db, c, blockType)

}

func (s *SQLStore) GetBoardAndCard(c store.Container, block *model.Block) (*model.Block, *model.Block, error) {
	return s.getBoardAndCard(s.db, c, block)

}

func (s *SQLStore) GetBoardAndCardByID(c store.Container, blockID string) (*model.Block, *model.Block, error) {
	return s.getBoardAndCardByID(s.db, c, blockID)

}

func (s *SQLStore) GetCardLimitTimestamp() (int64, error) {
	return s.getCardLimitTimestamp(s.db)

}

func (s *SQLStore) GetCloudLimits() (*mmModel.ProductLimits, error) {
	return s.getCloudLimits(s.db)

}

func (s *SQLStore) GetDefaultTemplateBlocks() ([]model.Block, error) {
	return s.getDefaultTemplateBlocks(s.db)

}

func (s *SQLStore) GetFileInfo(id string) (*mmModel.FileInfo, error) {
	return s.getFileInfo(s.db, id)

}

func (s *SQLStore) GetLicense() *mmModel.License {
	return s.getLicense(s.db)

}

func (s *SQLStore) GetNextNotificationHint(remove bool) (*model.NotificationHint, error) {
	return s.getNextNotificationHint(s.db, remove)

}

func (s *SQLStore) GetNotificationHint(c store.Container, blockID string) (*model.NotificationHint, error) {
	return s.getNotificationHint(s.db, c, blockID)

}

func (s *SQLStore) GetParentID(c store.Container, blockID string) (string, error) {
	return s.getParentID(s.db, c, blockID)

}

func (s *SQLStore) GetRegisteredUserCount() (int, error) {
	return s.getRegisteredUserCount(s.db)

}

func (s *SQLStore) GetRootID(c store.Container, blockID string) (string, error) {
	return s.getRootID(s.db, c, blockID)

}

func (s *SQLStore) GetSession(token string, expireTime int64) (*model.Session, error) {
	return s.getSession(s.db, token, expireTime)

}

func (s *SQLStore) GetSharing(c store.Container, rootID string) (*model.Sharing, error) {
	return s.getSharing(s.db, c, rootID)

}

func (s *SQLStore) GetSubTree2(c store.Container, blockID string, opts model.QuerySubtreeOptions) ([]model.Block, error) {
	return s.getSubTree2(s.db, c, blockID, opts)

}

func (s *SQLStore) GetSubTree3(c store.Container, blockID string, opts model.QuerySubtreeOptions) ([]model.Block, error) {
	return s.getSubTree3(s.db, c, blockID, opts)

}

func (s *SQLStore) GetSubscribersCountForBlock(c store.Container, blockID string) (int, error) {
	return s.getSubscribersCountForBlock(s.db, c, blockID)

}

func (s *SQLStore) GetSubscribersForBlock(c store.Container, blockID string) ([]*model.Subscriber, error) {
	return s.getSubscribersForBlock(s.db, c, blockID)

}

func (s *SQLStore) GetSubscription(c store.Container, blockID string, subscriberID string) (*model.Subscription, error) {
	return s.getSubscription(s.db, c, blockID, subscriberID)

}

func (s *SQLStore) GetSubscriptions(c store.Container, subscriberID string) ([]*model.Subscription, error) {
	return s.getSubscriptions(s.db, c, subscriberID)

}

func (s *SQLStore) GetSystemSetting(key string) (string, error) {
	return s.getSystemSetting(s.db, key)

}

func (s *SQLStore) GetSystemSettings() (map[string]string, error) {
	return s.getSystemSettings(s.db)

}

func (s *SQLStore) GetTeamBoardsInsights(duration string, channelIDs []string) ([]*model.BoardInsight, error) {
	return s.getTeamBoardsInsights(s.db, duration, channelIDs)

}

func (s *SQLStore) GetUserBoardsInsights(userID string, duration string, channelIDs []string) ([]*model.BoardInsight, error) {
	return s.getUserBoardsInsights(s.db, userID, duration, channelIDs)
}

func (s *SQLStore) GetUsedCardsCount() (int, error) {
	return s.getUsedCardsCount(s.db)

}

func (s *SQLStore) GetUserByEmail(email string) (*model.User, error) {
	return s.getUserByEmail(s.db, email)

}

func (s *SQLStore) GetUserByID(userID string) (*model.User, error) {
	return s.getUserByID(s.db, userID)

}

func (s *SQLStore) GetUserByUsername(username string) (*model.User, error) {
	return s.getUserByUsername(s.db, username)

}

func (s *SQLStore) GetUserWorkspaces(userID string) ([]model.UserWorkspace, error) {
	return s.getUserWorkspaces(s.db, userID)

}

func (s *SQLStore) GetUserWorkspacesInTeam(userID string, teamID string) ([]model.UserWorkspace, error) {
	return s.getUserWorkspacesInTeam(s.db, userID, teamID)

}

func (s *SQLStore) GetUsersByWorkspace(workspaceID string) ([]*model.User, error) {
	return s.getUsersByWorkspace(s.db, workspaceID)

}

func (s *SQLStore) GetWorkspace(ID string) (*model.Workspace, error) {
	return s.getWorkspace(s.db, ID)

}

func (s *SQLStore) GetWorkspaceCount() (int64, error) {
	return s.getWorkspaceCount(s.db)

}

func (s *SQLStore) GetWorkspaceTeam(workspaceID string) (*mmModel.Team, error) {
	return s.getWorkspaceTeam(s.db, workspaceID)

}

func (s *SQLStore) HasWorkspaceAccess(userID string, workspaceID string) (bool, error) {
	return s.hasWorkspaceAccess(s.db, userID, workspaceID)

}

func (s *SQLStore) InsertBlock(c store.Container, block *model.Block, userID string) error {
	if s.dbType == sqliteDBType {
		return s.insertBlock(s.db, c, block, userID)
	}
	tx, txErr := s.db.BeginTx(context.Background(), nil)
	if txErr != nil {
		return txErr
	}
	err := s.insertBlock(tx, c, block, userID)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			s.logger.Error("transaction rollback error", mlog.Err(rollbackErr), mlog.String("methodName", "InsertBlock"))
		}
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil

}

func (s *SQLStore) InsertBlocks(c store.Container, blocks []model.Block, userID string) error {
	if s.dbType == sqliteDBType {
		return s.insertBlocks(s.db, c, blocks, userID)
	}
	tx, txErr := s.db.BeginTx(context.Background(), nil)
	if txErr != nil {
		return txErr
	}
	err := s.insertBlocks(tx, c, blocks, userID)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			s.logger.Error("transaction rollback error", mlog.Err(rollbackErr), mlog.String("methodName", "InsertBlocks"))
		}
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil

}

func (s *SQLStore) PatchBlock(c store.Container, blockID string, blockPatch *model.BlockPatch, userID string) error {
	if s.dbType == sqliteDBType {
		return s.patchBlock(s.db, c, blockID, blockPatch, userID)
	}
	tx, txErr := s.db.BeginTx(context.Background(), nil)
	if txErr != nil {
		return txErr
	}
	err := s.patchBlock(tx, c, blockID, blockPatch, userID)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			s.logger.Error("transaction rollback error", mlog.Err(rollbackErr), mlog.String("methodName", "PatchBlock"))
		}
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil

}

func (s *SQLStore) PatchBlocks(c store.Container, blockPatches *model.BlockPatchBatch, userID string) error {
	if s.dbType == sqliteDBType {
		return s.patchBlocks(s.db, c, blockPatches, userID)
	}
	tx, txErr := s.db.BeginTx(context.Background(), nil)
	if txErr != nil {
		return txErr
	}
	err := s.patchBlocks(tx, c, blockPatches, userID)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			s.logger.Error("transaction rollback error", mlog.Err(rollbackErr), mlog.String("methodName", "PatchBlocks"))
		}
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil

}

func (s *SQLStore) PatchUserProps(userID string, patch model.UserPropPatch) error {
	return s.patchUserProps(s.db, userID, patch)

}

func (s *SQLStore) RefreshSession(session *model.Session) error {
	return s.refreshSession(s.db, session)

}

func (s *SQLStore) RemoveDefaultTemplates(blocks []model.Block) error {
	return s.removeDefaultTemplates(s.db, blocks)

}

func (s *SQLStore) SendMessage(message string, postType string, receipts []string) error {
	return s.sendMessage(s.db, message, postType, receipts)
}

//nolint:typecheck
func (s *SQLStore) SaveFileInfo(fileInfo *mmModel.FileInfo) error {
	return s.saveFileInfo(s.db, fileInfo)
}

func (s *SQLStore) SetSystemSetting(key string, value string) error {
	return s.setSystemSetting(s.db, key, value)

}

func (s *SQLStore) UndeleteBlock(c store.Container, blockID string, modifiedBy string) error {
	if s.dbType == sqliteDBType {
		return s.undeleteBlock(s.db, c, blockID, modifiedBy)
	}
	tx, txErr := s.db.BeginTx(context.Background(), nil)
	if txErr != nil {
		return txErr
	}
	err := s.undeleteBlock(tx, c, blockID, modifiedBy)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			s.logger.Error("transaction rollback error", mlog.Err(rollbackErr), mlog.String("methodName", "UndeleteBlock"))
		}
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil

}

func (s *SQLStore) UpdateCardLimitTimestamp(cardLimit int) (int64, error) {
	return s.updateCardLimitTimestamp(s.db, cardLimit)

}

func (s *SQLStore) UpdateSession(session *model.Session) error {
	return s.updateSession(s.db, session)

}

func (s *SQLStore) UpdateSubscribersNotifiedAt(c store.Container, blockID string, notifiedAt int64) error {
	return s.updateSubscribersNotifiedAt(s.db, c, blockID, notifiedAt)

}

func (s *SQLStore) UpdateUser(user *model.User) error {
	return s.updateUser(s.db, user)

}

func (s *SQLStore) UpdateUserPassword(username string, password string) error {
	return s.updateUserPassword(s.db, username, password)

}

func (s *SQLStore) UpdateUserPasswordByID(userID string, password string) error {
	return s.updateUserPasswordByID(s.db, userID, password)

}

func (s *SQLStore) UpsertNotificationHint(hint *model.NotificationHint, notificationFreq time.Duration) (*model.NotificationHint, error) {
	return s.upsertNotificationHint(s.db, hint, notificationFreq)

}

func (s *SQLStore) UpsertSharing(c store.Container, sharing model.Sharing) error {
	return s.upsertSharing(s.db, c, sharing)

}

func (s *SQLStore) UpsertWorkspaceSettings(workspace model.Workspace) error {
	return s.upsertWorkspaceSettings(s.db, workspace)

}

func (s *SQLStore) UpsertWorkspaceSignupToken(workspace model.Workspace) error {
	return s.upsertWorkspaceSignupToken(s.db, workspace)

}
