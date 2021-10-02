package app

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

func TestNewClient(t *testing.T) {
	testCases := []struct {
		name          string
		mnemomonicStr string
		expErr        bool
	}{
		{"success", Mnemonic, false},
		{"invalid mnemonic", "error", true},
		{"empty mnemonic", "", true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := NewClient(tc.mnemomonicStr)
			if tc.expErr {
				require.Error(t, err)
				require.Nil(t, res)
			} else {
				require.NoError(t, err)
				require.NotNil(t, res)
			}
		})
	}
}

func TestInitEthClient(t *testing.T) {
	testCases := []struct {
		name        string
		endpointStr string
		expErr      bool
	}{
		{"success", rpcEndpoint, false},
		{"invalid endpoint", "127.0.0.1:4546", true},
		{"empty endpoint", "", true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := initEthClient(tc.endpointStr)
			if tc.expErr {
				require.Error(t, err)
				require.Nil(t, res)
			} else {
				require.NoError(t, err)
				require.NotNil(t, res)
			}
		})
	}
}

type IntegrationTestSuite struct {
	suite.Suite

	client *Client
}

func (s *IntegrationTestSuite) SetupSuite() {
	client, err := NewClient(Mnemonic)
	s.Require().NoError(err)

	s.client = client
}

func (s *IntegrationTestSuite) TestDeployContract() {

	ethClient, err := initEthClient(rpcEndpoint)
	s.Require().NoError(err)

	privKey, err := privKeyFromMnemonic(Mnemonic)
	s.Require().NoError(err)

	testCases := []struct {
		name    string
		pretest func()
		expErr  bool
	}{
		{"privateKey is empty",
			func() {
				s.client.privateKey = nil
			},
			true,
		},
		{
			"ethClient is empty",
			func() {
				s.client.ethClient = nil
			},
			true,
		},
		{"success",
			func() {
				s.client = &Client{ethClient, privKey}
			},
			false},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			s.SetupSuite()
			tc.pretest()

			err := s.client.DeployContract()
			if tc.expErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)
			}
		})
	}
}

// TODO debug this test
/*
func (s *IntegrationTestSuite) TestGetContractInstance(){

	testCases := []struct {
		name        string
		contractHexStr string
		expErr      bool
	}{
		{"success", ContractAddr, false},
		{"invalid contract addr", ContractAddr + "bc", true},
		{"empty contract addr", "", true},
	}


	for _, tc := range testCases {
		s.Run(tc.name, func() {
			s.SetupSuite()

			res, err := s.client.GetContractInstance(tc.contractHexStr)
			if tc.expErr {
				if err != nil { s.T().Log(err.Error()) }
				s.Require().Error(err)
				s.Require().Nil(res)
			} else {
				s.Require().NoError(err)
				s.Require().NotNil(res)
			}
		})

	}
}
*/

// TODO
func TestIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}
