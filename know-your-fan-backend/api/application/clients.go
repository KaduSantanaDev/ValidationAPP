package application

import (
	"github.com/google/uuid"
)

type Client struct {
	ID       string
	Name     string
	Email    string
	CPF      string
	Document []byte
	Address  Address
	Status   bool
}

type ClientInterface interface {
	GetID() string
	GetName() string
	GetCPF() string
	GetEmail() string
	GetDocument() []byte
	GetAddress() Address
	GetStatus() bool
	UpdateStatus(status bool)
}

type ClientServiceInterface interface {
	Create(clientInterface ClientInterface) (ClientInterface, error)
}

func (c *Client) GetID() string {
	return c.ID
}
func (c *Client) GetName() string {
	return c.Name
}
func (c *Client) GetEmail() string {
	return c.Email
}
func (c *Client) GetCPF() string {
	return c.CPF
}
func (c *Client) GetDocument() []byte {
	return c.Document
}
func (c *Client) GetAddress() Address {
	return c.Address
}
func (c *Client) GetStatus() bool {
	return c.Status
}
func (c *Client) UpdateStatus(status bool) {
	c.Status = status
}

func NewClient() *Client {
	return &Client{
		ID: uuid.New().String(),
	}
}
