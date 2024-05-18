package db

import (
	"context"
	"database/sql"
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetAllSubscriptions(t *testing.T) {
	subscriptionNumber := 5
	for i := 0; i < subscriptionNumber; i++ {
		populateDBWithValidRandomSubscription(t)
	}

	subscriptions, err := testQueries.GetSubscriptions(context.Background())

	require.NoError(t, err)
	require.NotEmpty(t, subscriptions)
	require.Equal(t, subscriptionNumber, len(subscriptions))
}

func TestCreateSubscription(t *testing.T) {
	populateDBWithValidRandomSubscription(t)
}

func TestFailToCreateSubscriptionWithExistingEmail(t *testing.T) {
	subscription := populateDBWithValidRandomSubscription(t)

	subscription, err := testQueries.CreateSubscription(context.Background(), subscription.Email)

	require.Empty(t, subscription)
	require.NotNil(t, err)
	require.Equal(t, err.Error(), "pq: duplicate key value violates unique constraint \"subscriptions_email_key\"")
}

func TestGetCreatedSubscriptionByEmail(t *testing.T) {
	subscription := populateDBWithValidRandomSubscription(t)

	subscription, err := testQueries.GetSubscriptionByEmail(context.Background(), subscription.Email)

	checkInsertedSubscriptionIsValid(t, err, subscription, subscription.Email.(string))
}

func TestFailToGetNotExistingSubscriptionByEmail(t *testing.T) {
	email := faker.Email()

	subscription, err := testQueries.GetSubscriptionByEmail(context.Background(), email)

	require.Empty(t, subscription)
	require.ErrorIs(t, err, sql.ErrNoRows)
}

func populateDBWithValidRandomSubscription(t *testing.T) Subscription {
	email := faker.Email()

	subscription, err := testQueries.CreateSubscription(context.Background(), email)

	checkInsertedSubscriptionIsValid(t, err, subscription, email)

	return subscription
}

func checkInsertedSubscriptionIsValid(t *testing.T, err error, actual Subscription, email string) {
	require.NoError(t, err)
	require.NotEmpty(t, actual)

	require.NotZero(t, actual.ID)

	require.Equal(t, actual.Email.(string), email)
}
