
-- Create branch_transaction table
create table branch_transaction
(
  sysno        	number(20) not null CONSTRAINT branch_sysno_pk PRIMARY KEY,
  xid    	   	varchar2(128) not NULL UNIQUE ,
  branch_id     number(20) not null,
  args_json     varchar2(512) default null,
  state    		number(4) default null,
  gmt_create  	date DEFAULT SYSDATE, 
  gmt_modified 	DATE DEFAULT SYSDATE,
  CONSTRAINT xid_branchid_uk UNIQUE (xid, branch_id) 
);

CREATE SEQUENCE branch_transaction_seq START WITH 1 INCREMENT BY 1;

-- Create so_item table
CREATE TABLE so_item
(
  sysno        		NUMBER(20) NOT NULL CONSTRAINT item_sysno_pk PRIMARY KEY,
  so_sysno     		NUMBER(20) DEFAULT NULL,
  product_sysno     NUMBER(20) NOT NULL,
  product_name    	VARCHAR2(64) DEFAULT NULL,
  cost_price        NUMBER(16,6) DEFAULT NULL,
  original_price    NUMBER(16,6) DEFAULT NULL, 
  deal_price 		NUMBER(16,6) DEFAULT NULL,
  quantity 			NUMBER(11) DEFAULT NULL
);

CREATE SEQUENCE so_item_seq START WITH 1 INCREMENT BY 1;

-- Create so_master table
CREATE TABLE so_master (
  sysno 					NUMBER(20) NOT NULL CONSTRAINT master_sysno_pk PRIMARY KEY,
  so_id 					VARCHAR2(20) DEFAULT NULL,
  buyer_user_sysno 			NUMBER(20) DEFAULT NULL,
  seller_company_code  		VARCHAR2(20) DEFAULT NULL,
  receive_division_sysno 	NUMBER(20) NOT NULL,
  receive_address	 		VARCHAR2(200) DEFAULT NULL,
  receive_zip 				VARCHAR2(20) DEFAULT NULL,
  receive_contact 			VARCHAR2(20) DEFAULT NULL,
  receive_contact_phone 	VARCHAR2(100) DEFAULT NULL,
  stock_sysno				NUMBER(20) DEFAULT NULL,
  payment_type				NUMBER(4) DEFAULT NULL,
  so_amt 					NUMBER(16,6) DEFAULT NULL,
  status 					NUMBER(4) DEFAULT NULL,
  order_date				DATE DEFAULT NULL,
  payment_date				DATE DEFAULT NULL,
  delivery_date	 			DATE DEFAULT NULL,
  receive_date				DATE DEFAULT NULL,
  appid						VARCHAR2(64) DEFAULT NULL,
  memo						VARCHAR2(255) DEFAULT NULL,
  create_user 				VARCHAR2(255) DEFAULT NULL,
  gmt_create 				DATE DEFAULT NULL,
  modify_user 				VARCHAR2(255) DEFAULT NULL,
  gmt_modified 				DATE DEFAULT NULL
);

CREATE SEQUENCE so_master_seq START WITH 1 INCREMENT BY 1;

CREATE TABLE undo_log
(
    id            NUMBER(19)    NOT NULL CONSTRAINT undolog_sysno_pk PRIMARY KEY,
    branch_id     NUMBER(19)    NOT NULL,
    xid           VARCHAR2(128) NOT NULL,
    context       VARCHAR2(128) NOT NULL,
    rollback_info BLOB          NOT NULL,
    log_status    NUMBER(10)    NOT NULL,
    log_created   TIMESTAMP(0)  NOT NULL,
    log_modified  TIMESTAMP(0)  NOT NULL,
    CONSTRAINT ux_undo_log UNIQUE (xid, branch_id)
);

COMMENT ON TABLE undo_log IS 'AT transaction mode undo table';

-- Generate ID using sequence and trigger
CREATE SEQUENCE undo_log_seq START WITH 1 INCREMENT BY 1;

