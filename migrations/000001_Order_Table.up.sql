CREATE TABLE IF NOT EXISTS orders (
    OrderID SERIAL PRIMARY KEY,
	AuthorID INTEGER NOT NULL,
	DeliveryPrice INTEGER NOT NULL,
	ItemCategory VARCHAR(30) NOT NULL,
	ItemWeight VARCHAR(30) NOT NULL,
	FirstAddressPhone VARCHAR(15) NOT NULL,
	SecondAddressPhone VARCHAR(15) NOT NULL,
	CreatedAt DATE NOT NULL,
	FirstAddress JSON NOT NULL,
	SecondAddress JSON NOT NULL
);
