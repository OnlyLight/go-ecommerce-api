# Add connector
curl -i -X POST -H "Accept:application/json" -H "Content-Type:application/json" http://localhost:8083/connectors/ -d @config/connector/mysql-source.json

# # Check connector added
# curl -s localhost:8083/connector-plugins | jq '.[].class'
