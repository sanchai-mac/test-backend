export db ?= users_db

test:
	$(eval db=test_customer_db)
	$(eval PROFILE=test)