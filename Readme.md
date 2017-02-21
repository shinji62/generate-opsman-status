

# Create Bosh User into Uaa



```
uaac target https://uaa.[your cf system domain] --skip-ssl-validation
uaac token client get admin -s [your admin-secret]
uaac client add generate-opsman-statuss-users \
      --secret [your_client_secret] \
      --authorized_grant_types client_credentials,refresh_token \
      --authorities opsman.user \
      --scope opsman.user
```
