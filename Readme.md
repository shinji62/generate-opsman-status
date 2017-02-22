

# Create OpsManagerUser into Uaa
```
uaac target https://uaa.[youropsmanagerip] --skip-ssl-validation
uaac token client get admin -s [your admin-secret]
uaac client add generate-opsman-statuss-users \
      --secret [your_client_secret] \
      --authorized_grant_types client_credentials,refresh_token \
      --authorities opsman.admin \
      --scope opsman.admin
```



#Generate Json file
```shell
go run main.go --opsman-endpoint https://opsmanager.domain \
  --skip-ssl-validation \
  --client-id-opsman client-id-foropsman \
  --client-secret-opsman "version-pages-users-secret12" \
  --pivnet-api-token pivent-api-token
```
