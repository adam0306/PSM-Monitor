#!/usr/bin/python
import requests
from requests.packages.urllib3.exceptions import InsecureRequestWarning
# Ignoring HTTPS warnings.
requests.packages.urllib3.disable_warnings(InsecureRequestWarning)
# Defining the world.

failures = {}
# Replace "git" with the url of a raw list of FQDN servers that you would like to be monitored.
psmservers = filter(None, requests.get('git').text.split('\n'))
for psm in psmservers:
    r = requests.get('https://' + psm + './psm/api/health', verify=true)
    if r.text != 'PASS':
        failures.update({psm: r.text})
# Uncomment the line below if you would like for the failure messages to be printed to the screen.
#print(failures)
