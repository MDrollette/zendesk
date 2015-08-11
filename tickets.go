package zendesk

import (
	"fmt"
	"net/url"

	"golang.org/x/net/context"
)

type Ticket struct {
	Id                  int           `json:"id,omitempty"`
	Url                 string        `json:"url,omitempty"`
	ExternalId          string        `json:"external_id,omitempty"`
	CreatedAt           string        `json:"created_at,omitempty"`
	UpdatedAt           string        `json:"updated_at,omitempty"`
	Type                string        `json:"type,omitempty"`
	Subject             string        `json:"subject,omitempty"`
	RawSubject          string        `json:"raw_subject,omitempty"`
	Description         string        `json:"description,omitempty"`
	Priority            string        `json:"priority,omitempty"`
	Status              string        `json:"status,omitempty"`
	Recipient           string        `json:"recipient,omitempty"`
	RequesterId         int           `json:"requester_id,omitempty"`
	SubmitterId         int           `json:"submitter_id,omitempty"`
	AssigneeId          int           `json:"assignee_id,omitempty"`
	OrganizationId      int           `json:"organization_id,omitempty"`
	GroupId             int           `json:"group_id,omitempty"`
	CollaboratorIds     []int         `json:"collaborator_ids,omitempty"`
	SharingAgreementIds []int         `json:"sharing_agreement_ids,omitempty"`
	ForumTopicId        int           `json:"forum_topic_id,omitempty"`
	ProblemId           int           `json:"problem_id,omitempty"`
	HasIncidents        bool          `json:"has_incidents,omitempty"`
	DueAt               string        `json:"due_at,omitempty"`
	Tags                []string      `json:"tags,omitempty"`
	CustomFields        []CustomField `json:"custom_fields,omitempty"`
}

type CustomField struct {
	Id    int         `json:"id"`
	Value interface{} `json:"value"`
}

func (cf *CustomField) StringVal() string {
	if s, ok := cf.Value.(string); ok {
		return s
	}

	return ""
}

func (cf *CustomField) IntVal() int64 {
	if i, ok := cf.Value.(int64); ok {
		return i
	}

	return int64(0)
}

func (cf *CustomField) BoolVal() bool {
	if b, ok := cf.Value.(bool); ok {
		return b
	}

	return false
}

type TicketApi struct {
	client  *Client
	context context.Context
}

func (api *TicketApi) WithContext(ctx context.Context) *TicketApi {
	return &TicketApi{
		client:  api.client,
		context: ctx,
	}
}

func (api *TicketApi) getTickets(path string, params *url.Values) ([]Ticket, error) {
	response := struct {
		Tickets []Ticket `json:"tickets"`
	}{}

	err := api.client.get(api.context, path, params, &response)
	if err != nil {
		return nil, err
	}
	return response.Tickets, nil
}

func (api *TicketApi) getTicket(path string, params *url.Values) (Ticket, error) {
	response := struct {
		Ticket Ticket `json:"ticket"`
	}{}

	err := api.client.get(api.context, path, params, &response)
	if err != nil {
		return Ticket{}, err
	}
	return response.Ticket, nil
}

func (api *TicketApi) List() ([]Ticket, error) {
	return api.getTickets("/api/v2/tickets.json", nil)
}

func (api *TicketApi) Show(id int) (Ticket, error) {
	path := fmt.Sprintf("/api/v2/tickets/%d.json", id)
	return api.getTicket(path, nil)
}
