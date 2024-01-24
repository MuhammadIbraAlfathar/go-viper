package go_viper

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestViper(t *testing.T) {
	config := viper.New()

	assert.NotNil(t, config)
}

func TestJSON(t *testing.T) {
	config := viper.New()

	config.SetConfigName("config")
	config.SetConfigType("json")
	config.AddConfigPath(".")

	//read config
	err := config.ReadInConfig()
	if err != nil {
		fmt.Println(err.Error())
	}
	assert.Nil(t, err)

	assert.Equal(t, "test", config.GetString("App.name"))
}
