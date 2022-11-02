/*
   @dev: Logan (Nam) Nguyen
   @course: SUNY Oswego - CSC 495 - Capstone
   @instructor: Professor Bastian Tenbergen
   @version: 1.0
*/

// @package
package dao

// @import
import (
	"Swyl/servers/swyl-club-ms/models"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

// @notice Root struct for other methods in sub-dao-impl
type SubDaoImpl struct {
	ctx			 		context.Context
	mongoCollection 	*mongo.Collection
}

// @dev Constructor
func SubDaoConstructor(ctx context.Context, mongoCollection *mongo.Collection) SubDao {
	return &SubDaoImpl{
		ctx: ctx,
		mongoCollection: mongoCollection,
	}
}

// @notice Method of SubDaoImpl struct
// 
// @dev Lets a user subscribe to a tier
// 
// @param subParam *models.Subscription
// 
// @return error
func (si *SubDaoImpl) Subscribe(subParam *models.Subscription) error{return nil}


// @notice Method of SubDaoImpl struct
//
// @dev Gets a subscription using subscription_id
// 
// @param subId *string
// 
// @return *models.Subscription
func (si *SubDaoImpl) GetSubscriptionAt(subId *string) (*models.Subscription, error){return nil, nil}


// @notice Method of SubDaoImpl struct
//
// @dev Gets all subscriptions onwed at tier_id and by club_owner
// 
// @param tierId *string
// 
// @param clubOwner *string
// 
// @return *[]models.Subscription
// 
// @return error
func (si *SubDaoImpl) GetAllSubsAt(tierId *string, clubOwner *string) (*[]models.Subscription, error) {return nil, nil}


// @notice Method of SubDaoImpl struct
//
// @dev Updates a subscription status
// 
// @param subId *string
// 
// @return error
func (si *SubDaoImpl) UpdateSubStatus(subIb *string) error {return nil}


// @notice Method of SubDaoImpl struct
//
// @dev Lets a subscriber unsubscribe a tier 
// 
// @param tierId *string
// 
// @param subId *string
// 
// @return error
func (si *SubDaoImpl) Unsubscribe(tierId *string, subId *string) error {return nil}