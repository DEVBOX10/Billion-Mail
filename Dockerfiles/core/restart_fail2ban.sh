#!/bin/bash
if [[ ${FAIL2BAN_INIT} == "y" ]]; then
    /usr/bin/fail2ban-client reload core-accesslimit
    /usr/bin/fail2ban-client reload roundcube-accesslimit
fi
