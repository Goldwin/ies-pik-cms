package people

import (
	"context"

	. "github.com/Goldwin/ies-pik-cms/pkg/common/commands"
	"github.com/Goldwin/ies-pik-cms/pkg/common/out"
	q "github.com/Goldwin/ies-pik-cms/pkg/common/queries"
	"github.com/Goldwin/ies-pik-cms/pkg/common/worker"
	"github.com/Goldwin/ies-pik-cms/pkg/people/commands"
	"github.com/Goldwin/ies-pik-cms/pkg/people/dto"
	"github.com/Goldwin/ies-pik-cms/pkg/people/queries"
)

type PeopleDataLayerComponent interface {
	CommandWorker() worker.UnitOfWork[commands.CommandContext]
	QueryWorker() worker.QueryWorker[queries.QueryContext]
}

type PeopleManagementComponent interface {
	ViewPerson(context.Context, queries.ViewPersonQuery, out.Output[queries.ViewPersonResult])
	SearchPerson(context.Context, queries.SearchPersonQuery, out.Output[queries.SearchPersonResult])
	AddPerson(context.Context, dto.Person, out.Output[dto.Person])
	UpdatePerson(context.Context, dto.Person, out.Output[dto.Person])
	AddHousehold(context.Context, dto.HouseHoldInput, out.Output[dto.Household])
	UpdateHousehold(context.Context, dto.HouseHoldInput, out.Output[dto.Household])
	ViewHouseholdByPerson(context.Context, queries.ViewHouseholdByPersonQuery, out.Output[queries.ViewHouseholdByPersonResult])
}

func PeopleManagementComponents(worker worker.UnitOfWork[commands.CommandContext]) PeopleManagementComponent {
	return &peopleManagementComponent{
		worker: worker,
	}
}

type peopleManagementComponent struct {
	worker      worker.UnitOfWork[commands.CommandContext]
	queryWorker worker.QueryWorker[queries.QueryContext]
}

// ViewHouseholdByPerson implements PeopleManagementComponent.
func (p *peopleManagementComponent) ViewHouseholdByPerson(ctx context.Context, input queries.ViewHouseholdByPersonQuery, output out.Output[queries.ViewHouseholdByPersonResult]) {
	result, err := p.queryWorker.Query(ctx).ViewHouseholdByPerson().Execute(input)
	if err != q.NoQueryError {
		output.OnError(out.ConvertQueryErrorDetail(err))
	} else {
		output.OnSuccess(result)
	}
}

// SearchPerson implements PeopleManagementComponent.
func (p *peopleManagementComponent) SearchPerson(ctx context.Context, input queries.SearchPersonQuery, output out.Output[queries.SearchPersonResult]) {
	result, err := p.queryWorker.Query(ctx).SearchPerson().Execute(input)
	if err != q.NoQueryError {
		output.OnError(out.ConvertQueryErrorDetail(err))
	} else {
		output.OnSuccess(result)
	}
}

// ViewPerson implements PeopleManagementComponent.
func (p *peopleManagementComponent) ViewPerson(ctx context.Context, input queries.ViewPersonQuery, output out.Output[queries.ViewPersonResult]) {
	result, err := p.queryWorker.Query(ctx).ViewPerson().Execute(input)
	if err != q.NoQueryError {
		output.OnError(out.ConvertQueryErrorDetail(err))
	} else {
		output.OnSuccess(result)
	}
}

// AddHousehold implements PeopleManagementComponent.
func (p *peopleManagementComponent) AddHousehold(ctx context.Context, input dto.HouseHoldInput, output out.Output[dto.Household]) {
	var result CommandExecutionResult[dto.Household]
	_ = p.worker.Execute(ctx, func(ctx commands.CommandContext) error {
		result = commands.AddHouseholdCommand{Input: input}.Execute(ctx)
		if result.Status != ExecutionStatusSuccess {
			return result.Error
		}
		return nil
	})
	if result.Status == ExecutionStatusSuccess {
		output.OnSuccess(result.Result)
	} else {
		output.OnError(out.ConvertCommandErrorDetail(result.Error))
	}
}

// AddPerson implements PeopleManagementComponent.
func (p *peopleManagementComponent) AddPerson(ctx context.Context, input dto.Person, output out.Output[dto.Person]) {
	var result CommandExecutionResult[dto.Person]
	_ = p.worker.Execute(ctx, func(ctx commands.CommandContext) error {
		result = commands.AddPersonCommand{Input: input}.Execute(ctx)
		if result.Status != ExecutionStatusSuccess {
			return result.Error
		}
		return nil
	})
	if result.Status == ExecutionStatusSuccess {
		output.OnSuccess(result.Result)
	} else {
		output.OnError(out.ConvertCommandErrorDetail(result.Error))
	}

}

// UpdateHousehold implements PeopleManagementComponent.
func (p *peopleManagementComponent) UpdateHousehold(ctx context.Context, input dto.HouseHoldInput, output out.Output[dto.Household]) {
	var result CommandExecutionResult[dto.Household]
	_ = p.worker.Execute(ctx, func(ctx commands.CommandContext) error {
		result = commands.UpdateHouseholdCommand{Input: input}.Execute(ctx)
		if result.Status != ExecutionStatusSuccess {
			return result.Error
		}
		return nil
	})
	if result.Status == ExecutionStatusSuccess {
		output.OnSuccess(result.Result)
	} else {
		output.OnError(out.ConvertCommandErrorDetail(result.Error))
	}
}

// UpdatePerson implements PeopleManagementComponent.
func (p *peopleManagementComponent) UpdatePerson(ctx context.Context, input dto.Person, output out.Output[dto.Person]) {
	var result CommandExecutionResult[dto.Person]
	_ = p.worker.Execute(ctx, func(ctx commands.CommandContext) error {
		result = commands.UpdatePersonCommand{Input: input}.Execute(ctx)
		if result.Status != ExecutionStatusSuccess {
			return result.Error
		}
		return nil
	})
	if result.Status == ExecutionStatusSuccess {
		output.OnSuccess(result.Result)
	} else {
		output.OnError(out.ConvertCommandErrorDetail(result.Error))
	}
}

func NewPeopleManagementComponent(data PeopleDataLayerComponent) PeopleManagementComponent {
	return &peopleManagementComponent{
		worker:      data.CommandWorker(),
		queryWorker: data.QueryWorker(),
	}
}
