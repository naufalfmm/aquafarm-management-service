package frTime

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMock(t *testing.T) {
	var (
		now = time.Now()
	)

	t.Run("If mock success, it will freeze the current time", func(t *testing.T) {
		Mock(now)
		defer ResetMock()

		assert.True(t, isFreeze)
		assert.Equal(t, &now, frozenTime)
	})

	t.Run("If mock has been done, the second call won't affect", func(t *testing.T) {
		Mock(now)
		defer ResetMock()
		Mock(time.Now())

		assert.True(t, isFreeze)
		assert.Equal(t, &now, frozenTime)
	})
}

func TestNow(t *testing.T) {
	var (
		now = time.Now()
	)

	t.Run("If freeze active, it will return the frozen time", func(t *testing.T) {
		Mock(now)
		defer ResetMock()

		checkNow := Now()

		assert.True(t, isFreeze)
		assert.Equal(t, &checkNow, frozenTime)
	})

	t.Run("If freeze unactive, it will not return the frozen time", func(t *testing.T) {
		Mock(now)
		ResetMock()

		checkNow := Now()

		assert.False(t, isFreeze)
		assert.NotEqual(t, &checkNow, now)
	})
}

func TestRestMock(t *testing.T) {
	var (
		now = time.Now()
	)

	t.Run("If mock has not been resetted, it will still freeze", func(t *testing.T) {
		Mock(now)
		defer ResetMock()

		assert.True(t, isFreeze)
		assert.Equal(t, &now, frozenTime)
	})

	t.Run("If mock has been resetted, it will not freeze", func(t *testing.T) {
		Mock(now)
		ResetMock()

		assert.False(t, isFreeze)
		assert.Nil(t, frozenTime)
	})

	t.Run("If mock has been resetted, the second reset call won't affect", func(t *testing.T) {
		Mock(now)
		ResetMock()
		ResetMock()

		assert.False(t, isFreeze)
		assert.Nil(t, frozenTime)
	})
}
