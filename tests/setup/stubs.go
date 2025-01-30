package setup

var (
	CreateAccountsTable = `CREATE TABLE accounts (
    id bigint UNSIGNED NOT NULL AUTO_INCREMENT,
    created_at datetime(3) DEFAULT NULL,
    updated_at datetime(3) DEFAULT NULL,
    deleted_at datetime(3) DEFAULT NULL,
    document_number varchar(20) NOT NULL UNIQUE,
    balance decimal(20,5) DEFAULT 0.0,
    PRIMARY KEY (id),
    INDEX idx_accounts_deleted_at (deleted_at)
	);`

	CreateTransactionsTable = `CREATE TABLE transactions (
    id bigint UNSIGNED NOT NULL AUTO_INCREMENT,
    created_at datetime(3) DEFAULT NULL,
    updated_at datetime(3) DEFAULT NULL,
    deleted_at datetime(3) DEFAULT NULL,
    account_id bigint UNSIGNED NOT NULL,
    operation_type int NOT NULL,
    amount decimal(20,5) NOT NULL,
    event_timestamp datetime NOT NULL,
    PRIMARY KEY (id),
    INDEX idx_transactions_deleted_at (deleted_at),
    INDEX idx_transactions_account_id (account_id),
    CONSTRAINT fk_transactions_account FOREIGN KEY (account_id) REFERENCES accounts (id) 
        ON UPDATE CASCADE 
        ON DELETE CASCADE
	);`
)
