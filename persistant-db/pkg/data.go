package pkg

var Employees []User = []User{
	{
		Name:    "John Doe",
		Age:     "30",
		Contact: "john.doe@example.com",
		Company: "ABC Corp",
		Address: Address{
			City:    "New York",
			State:   "NY",
			Country: "USA",
			Pincode: "10001",
		},
	},
	{
		Name:    "Jane Smith",
		Age:     "25",
		Contact: "jane.smith@example.com",
		Company: "XYZ Inc",
		Address: Address{
			City:    "Los Angeles",
			State:   "CA",
			Country: "USA",
			Pincode: "90001",
		},
	},
	{
		Name:    "Alice Johnson",
		Age:     "35",
		Contact: "alice.johnson@example.com",
		Company: "DEF Corp",
		Address: Address{
			City:    "Chicago",
			State:   "IL",
			Country: "USA",
			Pincode: "60601",
		},
	},
	{
		Name:    "Bob Williams",
		Age:     "40",
		Contact: "bob.williams@example.com",
		Company: "GHI Inc",
		Address: Address{
			City:    "Houston",
			State:   "TX",
			Country: "USA",
			Pincode: "77001",
		},
	},
	{
		Name:    "Emily Davis",
		Age:     "28",
		Contact: "emily.davis@example.com",
		Company: "JKL Corp",
		Address: Address{
			City:    "San Francisco",
			State:   "CA",
			Country: "USA",
			Pincode: "94101",
		},
	},
	{
		Name:    "Michael Wilson",
		Age:     "32",
		Contact: "michael.wilson@example.com",
		Company: "MNO Inc",
		Address: Address{
			City:    "Seattle",
			State:   "WA",
			Country: "USA",
			Pincode: "98101",
		},
	},
	{
		Name:    "Sophia Brown",
		Age:     "27",
		Contact: "sophia.brown@example.com",
		Company: "PQR Corp",
		Address: Address{
			City:    "Boston",
			State:   "MA",
			Country: "USA",
			Pincode: "02101",
		},
	},
	{
		Name:    "William Taylor",
		Age:     "45",
		Contact: "william.taylor@example.com",
		Company: "STU Inc",
		Address: Address{
			City:    "Dallas",
			State:   "TX",
			Country: "USA",
			Pincode: "75201",
		},
	},
	{
		Name:    "Olivia Anderson",
		Age:     "31",
		Contact: "olivia.anderson@example.com",
		Company: "VWX Corp",
		Address: Address{
			City:    "Miami",
			State:   "FL",
			Country: "USA",
			Pincode: "33101",
		},
	},
}

func GetEmployeesSlice() []User {
	return Employees
}
