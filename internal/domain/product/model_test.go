package product

import (
	"testing"
)

func TestProduct_ChangeTitle(t *testing.T) {
	tests := []struct {
		name      string
		initial   string
		newTitle  string
		wantErr   error
		wantTitle string
	}{
		{
			name:      "successful title change",
			initial:   "Super product (+99)",
			newTitle:  "Super product (+77)",
			wantErr:   nil,
			wantTitle: "Super product (+77)",
		},
		{
			name:      "title doesn't change - same value",
			initial:   "Super product (+99)",
			newTitle:  "Super product (+99)",
			wantErr:   ErrTitleDoesntChanged,
			wantTitle: "Super product (+99)",
		},
		{
			name:      "title too long",
			initial:   "Super product (+99)",
			newTitle:  "------------------------------------------Super product",
			wantErr:   ErrTitleWrongFormat,
			wantTitle: "Super product (+99)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Product{title: tt.initial}

			err := p.ChangeTitle(tt.newTitle)

			if err != tt.wantErr {
				t.Errorf("ChangeTitle() error = %v, wantErr %v", err, tt.wantErr)
			}

			if p.title != tt.wantTitle {
				t.Errorf("ChangeTitle() title = %v, want %v", p.title, tt.wantTitle)
			}
		})
	}
}

func TestProduct_ChangeDescription(t *testing.T) {
	tests := []struct {
		name            string
		initial         string
		newDescription  string
		wantErr         error
		wantDescription string
	}{
		{
			name:            "successful description change",
			initial:         "Super description",
			newDescription:  "Super mega description",
			wantErr:         nil,
			wantDescription: "Super mega description",
		},
		{
			name:            "description doesn't change - same value",
			initial:         "Super description",
			newDescription:  "Super description",
			wantErr:         ErrDescriptionDoesntChanged,
			wantDescription: "Super description",
		},
		{
			name:            "description too long",
			initial:         "Super description",
			newDescription:  "--------------------------------------------------------------------Super product---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------",
			wantErr:         ErrDescriptionWrongFormat,
			wantDescription: "Super description",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Product{description: tt.initial}

			err := p.ChangeDescription(tt.newDescription)

			if err != tt.wantErr {
				t.Errorf("ChangeDescription() error = %v, wantErr %v", err, tt.wantErr)
			}

			if p.description != tt.wantDescription {
				t.Errorf("ChangeDescription() description = %v, want %v", p.description, tt.wantDescription)
			}
		})
	}
}

func TestProduct_ChangeType(t *testing.T) {
	tests := []struct {
		name     string
		initial  string
		newType  string
		wantErr  error
		wantType string
	}{
		{
			name:     "successful type change",
			initial:  "Tech",
			newType:  "Medical",
			wantErr:  nil,
			wantType: "Medical",
		},
		{
			name:     "type doesn't change - same value",
			initial:  "Tech",
			newType:  "Tech",
			wantErr:  ErrTypeDoesntChanged,
			wantType: "Tech",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Product{productType: tt.initial}

			err := p.ChangeType(tt.newType)

			if err != tt.wantErr {
				t.Errorf("ChangeType() error = %v, wantErr %v", err, tt.wantErr)
			}

			if p.productType != tt.wantType {
				t.Errorf("ChangeType() productType = %v, want %v", p.productType, tt.wantType)
			}
		})
	}
}

func TestProduct_ChangePrice(t *testing.T) {
	tests := []struct {
		name      string
		initial   float64
		newPrice  float64
		wantErr   error
		wantPrice float64
	}{
		{
			name:      "successful price change",
			initial:   12.5,
			newPrice:  32.0,
			wantErr:   nil,
			wantPrice: 32.0,
		},
		{
			name:      "price doesn't change - same value",
			initial:   12.5,
			newPrice:  12.5,
			wantErr:   ErrPriceDoesntChanged,
			wantPrice: 12.5,
		},
		{
			name:      "change from zero to positive",
			initial:   0.0,
			newPrice:  15.0,
			wantErr:   nil,
			wantPrice: 15.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Product{price: tt.initial}

			err := p.ChangePrice(tt.newPrice)

			if err != tt.wantErr {
				t.Errorf("ChangePrice() error = %v, wantErr %v", err, tt.wantErr)
			}

			if p.price != tt.wantPrice {
				t.Errorf("ChangePrice() price = %v, want %v", p.price, tt.wantPrice)
			}
		})
	}
}
