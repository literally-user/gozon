package order

import "testing"

func TestOrder_ChangeAddress(t *testing.T) {
	tests := []struct {
		name        string
		initial     string
		newAddress  string
		wantErr     error
		wantAddress string
	}{
		{
			name:        "successful address change",
			initial:     "Old Street 123",
			newAddress:  "New Street 456",
			wantErr:     nil,
			wantAddress: "New Street 456",
		},
		{
			name:        "address doesn't change - same value",
			initial:     "Same Street 789",
			newAddress:  "Same Street 789",
			wantErr:     ErrAddressDoesntChanged,
			wantAddress: "Same Street 789",
		},
		{
			name:        "change from empty to non-empty",
			initial:     "",
			newAddress:  "New Street 111",
			wantErr:     nil,
			wantAddress: "New Street 111",
		},
		{
			name:        "change to empty address",
			initial:     "Old Street 222",
			newAddress:  "",
			wantErr:     nil,
			wantAddress: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Order{Address: tt.initial}

			err := o.ChangeAddress(tt.newAddress)

			if err != tt.wantErr {
				t.Errorf("ChangeAddress() error = %v, wantErr %v", err, tt.wantErr)
			}

			if o.Address != tt.wantAddress {
				t.Errorf("ChangeAddress() Address = %v, want %v", o.Address, tt.wantAddress)
			}
		})
	}
}

func TestOrder_MarkAsCompleted(t *testing.T) {
	tests := []struct {
		name          string
		initial       bool
		wantErr       error
		wantCompleted bool
	}{
		{
			name:          "successfully mark as completed",
			initial:       false,
			wantErr:       nil,
			wantCompleted: true,
		},
		{
			name:          "already completed - no change",
			initial:       true,
			wantErr:       ErrCompletedStateDoesntChanged,
			wantCompleted: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Order{Completed: tt.initial}

			err := o.MarkAsCompleted()

			if err != tt.wantErr {
				t.Errorf("MarkAsCompleted() error = %v, wantErr %v", err, tt.wantErr)
			}

			if o.Completed != tt.wantCompleted {
				t.Errorf("MarkAsCompleted() Completed = %v, want %v", o.Completed, tt.wantCompleted)
			}
		})
	}
}

func TestOrder_MarkAsCanceled(t *testing.T) {
	tests := []struct {
		name         string
		initial      bool
		wantErr      error
		wantCanceled bool
	}{
		{
			name:         "successfully mark as canceled",
			initial:      false,
			wantErr:      nil,
			wantCanceled: true,
		},
		{
			name:         "already canceled - no change",
			initial:      true,
			wantErr:      ErrCompletedStateDoesntChanged,
			wantCanceled: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Order{Canceled: tt.initial}

			err := o.MarkAsCanceled()

			if err != tt.wantErr {
				t.Errorf("MarkAsCanceled() error = %v, wantErr %v", err, tt.wantErr)
			}

			if o.Canceled != tt.wantCanceled {
				t.Errorf("MarkAsCanceled() Canceled = %v, want %v", o.Canceled, tt.wantCanceled)
			}
		})
	}
}

func TestOrder_MarkAsTaken(t *testing.T) {
	tests := []struct {
		name      string
		initial   bool
		wantErr   error
		wantTaken bool
	}{
		{
			name:      "successfully mark as taken",
			initial:   false,
			wantErr:   nil,
			wantTaken: true,
		},
		{
			name:      "already taken - no change",
			initial:   true,
			wantErr:   ErrTakenStateDoesntChanged,
			wantTaken: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &Order{Taken: tt.initial}

			err := o.MarkAsTaken()

			if err != tt.wantErr {
				t.Errorf("MarkAsTaken() error = %v, wantErr %v", err, tt.wantErr)
			}

			if o.Taken != tt.wantTaken {
				t.Errorf("MarkAsTaken() Taken = %v, want %v", o.Taken, tt.wantTaken)
			}
		})
	}
}
