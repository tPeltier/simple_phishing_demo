drop database if exists hospital_db;
create database hospital_db;
use hospital_db;

create table employees (
  id int not null auto_increment,
  firstname varchar(100) default null,
  lastname varchar(100) default null,
  password varchar(100) default null,
  position enum('doctor', 'nurse', 'admin', 'receptionist'),
  badge_status bool,
  ssn int default null,
  primary key(id)
) Engine=InnoDB;

create table patients (
  id int not null auto_increment,
  firstname varchar(100) default null,
  lastname varchar(100) default null,
  ssn int default null,
  ehr varchar(100) default null,
  primary key(id)
) Engine=InnoDB;

insert into employees
  (firstname, lastname, password, position, badge_status, ssn)
values
  ('John', 'Smith', 'password1', 'doctor', false, 12345),
  ('Jane', 'Doe', 'password2', 'nurse', false, 67890),
  ('Bob', 'Ryan', 'password3', 'receptionist', false, 13579),
  ('Alison', 'McDaniel', 'password1', 'admin', false, 24680);

insert into patients
  (firstname, lastname, ssn, ehr)
values
  ('Bruno', 'Holzer', 12457, 'Lab Results | Hemoglobin A1C result is 6.2%, consistent with pre-diabetes.'),
  ('Hale', 'Greenspan', 12457, 'Lab Results | Blood sugar level is high at 250 mg/dL, indicating hyperglycemia.'),
  ('Benedicte', 'Lyon', 98472, 'Lab Results | Sodium level is normal at 140 mEq/L.'),
  ('Linde', 'Geary', 34780, 'Lab Results | Liver enzymes are slightly elevated with ALT at 60 U/L and AST at 50 U/L.');
