package impl

import (
	"context"
	"testing"

	proto "{{ cookiecutter.repo_host }}/{{ cookiecutter.repo_owner }}/{{ cookiecutter.repo_name }}/proto"

	assert "github.com/stretchr/testify/assert"
	suite "github.com/stretchr/testify/suite"
	zap "go.uber.org/zap"
)

type TestSuite struct {
	suite.Suite
}

func TestRunSuit(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (suite *TestSuite) TestWithNameFalseNoNameProvided() {
	service := Service{
		Logger: zap.L(),
		Config: ServiceConfig{
			WithName: false,
		},
	}

	res, err := service.SayHello(context.Background(), &proto.HelloRequest{
		Name: "",
	})

	assert.Nil(suite.T(), err, "Error returned for valid request")
	assert.Equal(suite.T(), "Hello, World!", res.Message)
}

func (suite *TestSuite) TestWithNameFalseNameProvided() {
	service := Service{
		Logger: zap.L(),
		Config: ServiceConfig{
			WithName: false,
		},
	}

	res, err := service.SayHello(context.Background(), &proto.HelloRequest{
		Name: "",
	})

	assert.Nil(suite.T(), err, "Error returned for valid request")
	assert.Equal(suite.T(), "Hello, World!", res.Message)
}

func (suite *TestSuite) TestWithNameTrueNameProvided() {
	service := Service{
		Logger: zap.L(),
		Config: ServiceConfig{
			WithName: true,
		},
	}

	res, err := service.SayHello(context.Background(), &proto.HelloRequest{
		Name: "Some Name",
	})

	assert.Nil(suite.T(), err, "Error returned for valid request")
	assert.Equal(suite.T(), "Hello, Some Name!", res.Message)
}

func (suite *TestSuite) TestWithNameTrueNoNameProvided() {
	service := Service{
		Logger: zap.L(),
		Config: ServiceConfig{
			WithName: true,
		},
	}

	res, err := service.SayHello(context.Background(), &proto.HelloRequest{
		Name: "",
	})

	assert.Error(suite.T(), err, "Error must be returned")
	assert.Nil(suite.T(), res, "The response must be nil")
}
