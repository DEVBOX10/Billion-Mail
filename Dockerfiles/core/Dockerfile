FROM alpine:3.20

# Set environment variables
ENV LANG=C.UTF-8
ENV LC_ALL=C.UTF-8

# Copy file
# COPY repositories /etc/apk/repositories
COPY stop-supervisor.sh /stop-supervisor.sh
COPY core.sh /core.sh
COPY restart_fail2ban.sh /restart_fail2ban.sh

# Install dependencies (using Alpine's apk package manager)
RUN apk add --no-cache \
    bash \
    ca-certificates \
    curl \
    supervisor \
    rsyslog \
    tzdata \
    busybox-extras \
    postgresql-client \
    fail2ban \
    iptables \
    ipset \
    && rm -rf /var/cache/apk/* \
    && chmod +x /stop-supervisor.sh /core.sh /restart_fail2ban.sh

# Copy file

COPY supervisord.conf /etc/supervisor/supervisord.conf
COPY fail2ban.conf /etc/fail2ban/fail2ban.conf


ENTRYPOINT ["/core.sh"]
CMD ["/usr/bin/supervisord", "-c", "/etc/supervisor/supervisord.conf"]