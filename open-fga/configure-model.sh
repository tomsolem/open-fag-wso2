# bin/bash!

# Create a authorization model for the FGA Demo Store
# run the tests
 fga model test --tests /workspace/open-fga/store.yaml
# if all test pass, import the model to the store (database)
# Run the command and store its JSON output in a variable
FGA_IMPORT_OUTPUT=$(fga store import --file /workspace/open-fga/store.yaml)

# Use jq to parse the JSON and export the values to environment variables
export FGA_STORE_ID=$(echo "$FGA_IMPORT_OUTPUT" | jq -r '.store.id')
export FGA_MODEL_ID=$(echo "$FGA_IMPORT_OUTPUT" | jq -r '.model.authorization_model_id')

# Verify that the variables are set
echo "FGA_STORE_ID=${FGA_STORE_ID}"
echo "FGA_MODEL_ID=${FGA_MODEL_ID}" 