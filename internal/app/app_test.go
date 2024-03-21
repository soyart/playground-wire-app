package app

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"example.com/playground-wire-app/internal/config"
	"example.com/playground-wire-app/internal/domain/mock_repo"
	"example.com/playground-wire-app/internal/logger"
)

func TestApp(t *testing.T) {
	t.Run("app name should get result", func(t *testing.T) {
		appName := "appName"
		expectedData := fmt.Sprintf("Some data for key %s", appName)

		ctrl := gomock.NewController(t)
		mockRepo := mock_repo.NewMockRepo(ctrl)

		// Expect that the mockRepo.Read(appName) will only be called once,
		// and with arg appName only
		mockRepo.EXPECT().
			Read(gomock.Eq(appName)).
			Return([]byte(expectedData), nil)

		// Expect that mockRepo.Close() will be called only once
		mockRepo.EXPECT().Close()

		app := App{
			Configuration: config.Config{AppName: appName},
			Repository:    mockRepo,
			Logger:        new(logger.LoggerBasic),
		}

		err := app.Run()
		require.NoError(t, err)
	})
}
