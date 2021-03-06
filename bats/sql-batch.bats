#!/usr/bin/env bats
load $BATS_TEST_DIRNAME/helper/common.bash

setup() {
    setup_common
    dolt sql <<SQL
CREATE TABLE test (
  pk bigint NOT NULL,
  c1 bigint,
  c2 bigint,
  c3 bigint,
  c4 bigint,
  c5 bigint,
  PRIMARY KEY (pk)
);
SQL
}

teardown() {
    teardown_common
}

@test "dolt sql -b and -batch are a valid commands" {
    run dolt sql -b -q "insert into test values (0,0,0,0,0,0)"
    [ "$status" -eq 0 ]
    [[ "$output" =~ "Rows inserted: 1" ]] || false
    dolt sql -batch -q "insert into test values (1,0,0,0,0,0)"
    run dolt sql -b -q "select * from test" 
    [ "$status" -eq 0 ]
    [[ "$output" =~ " 0 " ]] || false
    [[ "$output" =~ " 1 " ]] || false

    # check for trailing newline
    # https://backreference.org/2010/05/23/sanitizing-files-with-no-trailing-newline/
    dolt sql -batch -q "insert into test values (2,0,0,0,0,0)" > out.txt
    lastline=$(tail -n 1 out.txt; echo x)
    lastline=${lastline%x}
    [[ $lastline =~ $'\n'$ ]] || false
}

@test "Piped SQL files interpreted in batch mode" {
    run dolt sql <<SQL
insert into test values (0,0,0,0,0,0);
insert into test values (1,0,0,0,0,0);
SQL
    [ "$status" -eq 0 ]
    [[ "$output" =~ "Rows inserted: 2" ]] || false
}

@test "Line number and bad query displayed on error in batch sql" {
    run dolt sql <<SQL
insert into test values (0,0,0,0,0,0);
insert into test values (1,0,0,0,0,0);
insert into test values poop;
SQL
    [ "$status" -ne 0 ]
    [[ "$output" =~ "Error processing batch" ]] || false
    [[ "$output" =~ "error on line 3 for query" ]] || false
    [[ "$output" =~ "insert into test values poop" ]] || false

    run dolt sql <<SQL
insert into test values (0,0,0,0,0,0);

insert into test values (1,0,
0,0,0,0);

insert into 
test values (2,0,0,0,0,0)
;

insert into 
test values 
poop;

insert into test values (3,0,0,0,0,0);
SQL
    [ "$status" -ne 0 ]
    [[ "$output" =~ "Error processing batch" ]] || false
    [[ "$output" =~ "error on line 10 for query" ]] || false
    [[ "$output" =~ "poop" ]] || false
}

@test "sql dolt_reset('hard') function" {
    mkdir test && cd test && dolt init
    dolt sql <<SQL
CREATE TABLE test (
    pk int PRIMARY KEY,
    c0 int
);
SQL
    dolt add -A && dolt commit -m "added table test"

    run dolt status
    [ "$status" -eq 0 ]
    [[ "$output" =~ "working tree clean" ]] || false

    dolt sql <<SQL
INSERT INTO test VALUES (1,1);
SQL
    run dolt status
    [ "$status" -eq 0 ]
    [[ "$output" =~ "test" ]] || false

   dolt sql <<SQL
SET @@test_head=dolt_reset('hard');
REPLACE INTO dolt_branches (hash,name) VALUES (@@test_head,'master');
SQL

    run dolt status
    [ "$status" -eq 0 ]
    [[ "$output" =~ "working tree clean" ]] || false

    dolt sql <<SQL
INSERT INTO test VALUES (1,1);
SET @@test_head = dolt_reset('hard');
REPLACE INTO dolt_branches (hash,name) VALUES (@@test_head,'master');
SQL

    run dolt status
    [ "$status" -eq 0 ]
    [[ "$output" =~ "working tree clean" ]] || false
}
