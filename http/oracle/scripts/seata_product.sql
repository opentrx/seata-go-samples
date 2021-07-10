-- Create branch_transaction table
create table branch_transaction
(
  sysno        	number(20) not null CONSTRAINT branch_sysno_pk PRIMARY KEY,
  xid    	   	varchar2(128) not NULL UNIQUE ,
  branch_id     number(20) not null,
  args_json     varchar2(512) default null,
  state    		number(4) default null,
  gmt_create  	date DEFAULT SYSDATE, 
  gmt_modified 	date DEFAULT SYSDATE,
  CONSTRAINT xid_branchid_uk UNIQUE (xid, branch_id) 
);

CREATE SEQUENCE branch_transaction_seq START WITH 1 INCREMENT BY 1;

CREATE TABLE inventory (
  sysno 				number(20) NOT NULL CONSTRAINT inventory_sysno_pk PRIMARY KEY,
  product_sysno 		number(20) NOT NULL,
  account_qty 		    number(11) DEFAULT NULL,
  available_qty 		number(11) DEFAULT NULL,
  allocated_qty 		number(11) DEFAULT NULL,
  adjust_locked_qty 	number(11) DEFAULT NULL
);


CREATE SEQUENCE inventory_seq START WITH 1 INCREMENT BY 1;

-- ----------------------------
-- Records of inventory
-- ----------------------------
BEGIN;
INSERT INTO inventory VALUES (1, 1, 1000000, 1000000, 0, 0);
COMMIT;

-- ----------------------------
-- Table structure for product
-- ----------------------------
CREATE TABLE product (
  sysno 			number(20) NOT NULL CONSTRAINT product_sysno_pk PRIMARY KEY,
  product_name 		varchar2(32) NOT NULL,
  product_title 	varchar2(32) NOT NULL,
  product_desc 		varchar2(2000) NOT NULL,
  product_desc_long clob NOT NULL,
  default_image_src varchar2(200) DEFAULT NULL,
  c3_sysno 			number(20) DEFAULT NULL,
  barcode 			varchar2(30) DEFAULT NULL,
  leng 				number(11) DEFAULT NULL,
  width 			number(11) DEFAULT NULL,
  height 			number(11) DEFAULT NULL,
  weight 			float DEFAULT NULL,
  merchant_sysno 	number(20) DEFAULT NULL,
  merchant_productid varchar2(20) DEFAULT NULL,
  status 			number(4) DEFAULT '1' NOT NULL,
  gmt_create  		date DEFAULT SYSDATE NOT NULL,
  create_user 		varchar2(32) NOT NULL,
  modify_user 		varchar2(32) NOT NULL,
  gmt_modified 		date DEFAULT SYSDATE NOT NULL
);

-- ----------------------------
-- Records of product
-- ----------------------------
BEGIN;
INSERT INTO product VALUES (1, '刺力王', '从小喝到大的刺力王', '好喝好喝好好喝', ' ', 'https://img10.360buyimg.com/mobilecms/s500x500_jfs/t1/61921/34/1166/131384/5cf60a94E411eee07/1ee010f4142236c3.jpg', 0, ' ', 15, 5, 5, 5, 1, ' ', 1, sysdate, ' ', ' ', sysdate);
COMMIT;

CREATE SEQUENCE product_seq START WITH 1 INCREMENT BY 1;

-- ----------------------------
-- If table exists, delete
-- ---------------------------
declare
      num   number;
begin
    select count(1) into num from user_tables where table_name = upper('undo_log') ;
    if num > 0 then
        execute immediate 'drop table undo_log' ;
    end if;
end;
-- ----------------------------
-- Table structure for undo_log
-- ---------------------------

CREATE TABLE undo_log
(
    id            NUMBER(19)    NOT NULL,
    branch_id     NUMBER(19)    NOT NULL,
    xid           VARCHAR2(128) NOT NULL,
    context       VARCHAR2(128) NOT NULL,
    rollback_info BLOB          NOT NULL,
    log_status    NUMBER(10)    NOT NULL,
    log_created   TIMESTAMP(0)  NOT NULL,
    log_modified  TIMESTAMP(0)  NOT NULL,
    PRIMARY KEY (id),
    CONSTRAINT ux_undo_log UNIQUE (xid, branch_id)
);

COMMENT ON TABLE undo_log IS 'AT transaction mode undo table';

-- Generate ID using sequence and trigger
CREATE SEQUENCE undo_log_seq START WITH 1 INCREMENT BY 1;

