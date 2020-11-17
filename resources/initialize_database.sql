CREATE TYPE DOCUMENT_TYPE AS ENUM('CPF', 'CNPJ');
CREATE TABLE accounts (
    id VARCHAR(36) PRIMARY KEY NOT NULL,
    document_type DOCUMENT_TYPE NOT NULL,
    document_number VARCHAR NOT NULL
);


CREATE TYPE OPERATION_FINALITY AS ENUM('CREDIT', 'DEBIT');
CREATE TABLE operations (
    id VARCHAR(36) PRIMARY KEY NOT NULL,
    description VARCHAR NOT NULL,
    finality OPERATION_FINALITY NOT NULL
);

CREATE TABLE transactions (
    id VARCHAR(36) PRIMARY KEY NOT NULL,
    operation_id VARCHAR NOT NULL,
    account_id VARCHAR NOT NULL,
    amount DECIMAL(10,2) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    CONSTRAINT fk_transactions_operation FOREIGN KEY (operation_id) REFERENCES operations (id),
    CONSTRAINT fk_transactions_account FOREIGN KEY (account_id) REFERENCES accounts (id)
);

INSERT INTO accounts (id, document_type, document_number) VALUES ('4e989872-6cc9-4746-a91c-8228afba818b', 'CPF', '39758803093');

INSERT INTO operations (id, description, finality) VALUES ('de8ba5ed-09b8-4555-9f0c-0e11323a4d62', 'COMPRA A VISTA', 'DEBIT');
INSERT INTO operations (id, description, finality) VALUES ('a39ebb08-2730-44c8-9a34-e2d51355b561', 'COMPRA PARCELADA', 'DEBIT');
INSERT INTO operations (id, description, finality) VALUES ('9217a7fe-5c40-4917-af2d-ff76a255b6d5', 'SAQUE', 'DEBIT');
INSERT INTO operations (id, description, finality) VALUES ('c030116a-8d26-4feb-ab91-6d1c025d7515', 'PAGAMENTO', 'CREDIT');

INSERT INTO transactions (id, operation_id, account_id, amount, created_at) VALUES ('38cc4090-1b33-4f11-a8e2-cdb3a9d260ef', 'de8ba5ed-09b8-4555-9f0c-0e11323a4d62', '4e989872-6cc9-4746-a91c-8228afba818b', 50.0, NOW());
INSERT INTO transactions (id, operation_id, account_id, amount, created_at) VALUES ('71cfe893-6cb5-40f8-a065-59a9ab39eb41', 'de8ba5ed-09b8-4555-9f0c-0e11323a4d62', '4e989872-6cc9-4746-a91c-8228afba818b', 23.5, NOW());
INSERT INTO transactions (id, operation_id, account_id, amount, created_at) VALUES ('edf12281-cf0b-421f-9480-b1fb3c999ac7', 'de8ba5ed-09b8-4555-9f0c-0e11323a4d62', '4e989872-6cc9-4746-a91c-8228afba818b', 18.7, NOW());
INSERT INTO transactions (id, operation_id, account_id, amount, created_at) VALUES ('4cebc9d2-f3e1-4d75-bcec-bb30b3afce00', 'c030116a-8d26-4feb-ab91-6d1c025d7515', '4e989872-6cc9-4746-a91c-8228afba818b', 60, NOW());