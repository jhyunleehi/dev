+functions-common:write_uwsgi_user_unit_file:1600  sudo systemctl daemon-reload
+functions-common:_run_under_systemd:1654  sudo systemctl enable devstack@keystone.service
Created symlink /etc/systemd/system/multi-user.target.wants/devstack@keystone.service → /etc/systemd/system/devstack@keystone.service.
+functions-common:_run_under_systemd:1655  sudo systemctl start devstack@keystone.service
+functions-common:run_process:1687         time_stop run_process
+functions-common:time_stop:2421           local name
+functions-common:time_stop:2422           local end_time
+functions-common:time_stop:2423           local elapsed_time
+functions-common:time_stop:2424           local total
+functions-common:time_stop:2425           local start_time
+functions-common:time_stop:2427           name=run_process
+functions-common:time_stop:2428           start_time=1714642712323
+functions-common:time_stop:2430           [[ -z 1714642712323 ]]
++functions-common:time_stop:2433           date +%s%3N
+functions-common:time_stop:2433           end_time=1714642713933
+functions-common:time_stop:2434           elapsed_time=1610
+functions-common:time_stop:2435           total=1563
+functions-common:time_stop:2437           _TIME_START[$name]=
+functions-common:time_stop:2438           _TIME_TOTAL[$name]=3173
+lib/keystone:start_keystone:569           echo 'Waiting for keystone to start...'
Waiting for keystone to start...
+lib/keystone:start_keystone:575           local service_uri=http://192.168.137.103/identity/v3/
+lib/keystone:start_keystone:577           wait_for_service 60 http://192.168.137.103/identity/v3/
+functions:wait_for_service:468            local timeout=60
+functions:wait_for_service:469            local url=http://192.168.137.103/identity/v3/
+functions:wait_for_service:470            local rval=0
+functions:wait_for_service:471            time_start wait_for_service
+functions-common:time_start:2407          local name=wait_for_service
+functions-common:time_start:2408          local start_time=
+functions-common:time_start:2409          [[ -n '' ]]
++functions-common:time_start:2412          date +%s%3N
+functions-common:time_start:2412          _TIME_START[$name]=1714642714032
+functions:wait_for_service:472            timeout 60 bash -x
++::                                        curl -g -k --noproxy '*' -s -o /dev/null -w '%{http_code}' http://192.168.137.103/identity/v3/
+::                                        [[ 200 == 503 ]]
+::                                        [[ 0 -eq 7 ]]
+functions:wait_for_service:477            time_stop wait_for_service
+functions-common:time_stop:2421           local name
+functions-common:time_stop:2422           local end_time
+functions-common:time_stop:2423           local elapsed_time
+functions-common:time_stop:2424           local total
+functions-common:time_stop:2425           local start_time
+functions-common:time_stop:2427           name=wait_for_service
+functions-common:time_stop:2428           start_time=1714642714032
+functions-common:time_stop:2430           [[ -z 1714642714032 ]]
++functions-common:time_stop:2433           date +%s%3N
+functions-common:time_stop:2433           end_time=1714642715997
+functions-common:time_stop:2434           elapsed_time=1965
+functions-common:time_stop:2435           total=0
+functions-common:time_stop:2437           _TIME_START[$name]=
+functions-common:time_stop:2438           _TIME_TOTAL[$name]=1965
+functions:wait_for_service:478            return 0
+lib/keystone:start_keystone:582           is_service_enabled tls-proxy
+functions-common:is_service_enabled:2048  local xtrace
++functions-common:is_service_enabled:2049  set +o
++functions-common:is_service_enabled:2049  grep xtrace
+functions-common:is_service_enabled:2049  xtrace='set -o xtrace'
+functions-common:is_service_enabled:2050  set +o xtrace
+functions-common:is_service_enabled:2077  return 1
+lib/keystone:start_keystone:587           restart_service memcached
+functions-common:restart_service:2305     '[' -x /bin/systemctl ']'
+functions-common:restart_service:2306     sudo /bin/systemctl restart memcached
+./stack.sh:main:1130                      bootstrap_keystone
+lib/keystone:bootstrap_keystone:607       /opt/stack/data/venv/bin/keystone-manage bootstrap --bootstrap-username admin --bootstrap-password 1234qwer --bootstrap-project-name admin --bootstrap-role-name admin --bootstrap-service-name keystone --bootstrap-region-id RegionOne --bootstrap-public-url http://192.168.137.103/identity
DEBUG keystone.notifications [-] Callback: `keystone.application_credential.core.Manager._delete_app_creds_on_user_delete_callback` subscribed to event `identity.user.deleted`. {{(pid=222896) register_event_callback /opt/stack/keystone/keystone/notifications.py:295}}
DEBUG keystone.notifications [-] Callback: `keystone.application_credential.core.Manager._delete_app_creds_on_user_delete_callback` subscribed to event `identity.user.disabled`. {{(pid=222896) register_event_callback /opt/stack/keystone/keystone/notifications.py:295}}
DEBUG keystone.notifications [-] Callback: `keystone.application_credential.core.Manager._delete_app_creds_on_assignment_removal` subscribed to event `identity.remove_application_credentials_for_user.internal`. {{(pid=222896) register_event_callback /opt/stack/keystone/keystone/notifications.py:295}}
DEBUG keystone.notifications [-] Callback: `keystone.assignment.core.Manager._delete_domain_assignments` subscribed to event `identity.domain.deleted`. {{(pid=222896) register_event_callback /opt/stack/keystone/keystone/notifications.py:295}}
DEBUG keystone.notifications [-] Callback: `keystone.catalog.core.Manager._on_project_or_endpoint_delete` subscribed to event `identity.project.deleted`. {{(pid=222896) register_event_callback /opt/stack/keystone/keystone/notifications.py:295}}
DEBUG keystone.notifications [-] Callback: `keystone.catalog.core.Manager._on_project_or_endpoint_delete` subscribed to event `identity.endpoint.deleted`. {{(pid=222896) register_event_callback /opt/stack/keystone/keystone/notifications.py:295}}
DEBUG keystone.notifications [-] Callback: `keystone.federation.core.Manager._cleanup_identity_provider` subscribed to event `identity.domain_deleted.internal`. {{(pid=222896) register_event_callback /opt/stack/keystone/keystone/notifications.py:295}}
DEBUG keystone.notifications [-] Callback: `keystone.identity.core.Manager._domain_deleted` subscribed to event `identity.domain_deleted.internal`. {{(pid=222896) register_event_callback /opt/stack/keystone/keystone/notifications.py:295}}
DEBUG keystone.notifications [-] Callback: `keystone.identity.core.Manager._unset_default_project` subscribed to event `identity.project.deleted`. {{(pid=222896) register_event_callback /opt/stack/keystone/keystone/notifications.py:295}}
DEBUG keystone.notifications [-] Callback: `keystone.revoke.core.Manager._trust_callback` subscribed to event `identity.OS-TRUST:trust.deleted`. {{(pid=222896) register_event_callback /opt/stack/keystone/keystone/notifications.py:295}}
DEBUG keystone.notifications [-] Callback: `keystone.revoke.core.Manager._consumer_callback` subscribed to event `identity.OS-OAUTH1:consumer.deleted`. {{(pid=222896) register_event_callback /opt/stack/keystone/keystone/notifications.py:295}}
DEBUG keystone.notifications [-] Callback: `keystone.revoke.core.Manager._user_callback` subscribed to event `identity.user.deleted`. {{(pid=222896) register_event_callback /opt/stack/keystone/keystone/notifications.py:295}}
DEBUG keystone.notifications [-] Callback: `keystone.revoke.core.Manager._project_callback` subscribed to event `identity.project.deleted`. {{(pid=222896) register_event_callback /opt/stack/keystone/keystone/notifications.py:295}}
DEBUG keystone.notifications [-] Callback: `keystone.revoke.core.Manager._user_callback` subscribed to event `identity.user.disabled`. {{(pid=222896) register_event_callback /opt/stack/keystone/keystone/notifications.py:295}}
DEBUG keystone.notifications [-] Callback: `keystone.revoke.core.Manager._user_callback` subscribed to event `identity.persist_revocation_event_for_user.internal`. {{(pid=222896) register_event_callback /opt/stack/keystone/keystone/notifications.py:295}}
DEBUG keystone.notifications [-] Callback: `keystone.receipt.provider.Manager._drop_receipt_cache` subscribed to event `identity.OS-TRUST:trust.deleted`. {{(pid=222896) register_event_callback /opt/stack/keystone/keystone/notifications.py:295}}
DEBUG keystone.notifications [-] Callback: `keystone.receipt.provider.Manager._drop_receipt_cache` subscribed to event `identity.user.deleted`. {{(pid=222896) register_event_callback /opt/stack/keystone/keystone/notifications.py:295}}
DEBUG keystone.notifications [-] Callback: `keystone.receipt.provider.Manager._drop_receipt_cache` subscribed to event `identity.domain.deleted`. {{(pid=222896) register_event_callback /opt/stack/keystone/keystone/notifications.py:295}}
DEBUG keystone.notifications [-] Callback: `keystone.receipt.provider.Manager._drop_receipt_cache` subscribed to event `identity.user.disabled`. {{(pid=222896) register_event_callback /opt/stack/keystone/keystone/notifications.py:295}}
DEBUG keystone.notifications [-] Callback: `keystone.receipt.provider.Manager._drop_receipt_cache` subscribed to event `identity.domain.disabled`. {{(pid=222896) register_event_callback /opt/stack/keystone/keystone/notifications.py:295}}
DEBUG keystone.notifications [-] Callback: `keystone.receipt.provider.Manager._drop_receipt_cache` subscribed to event `identity.project.disabled`. {{(pid=222896) register_event_callback /opt/stack/keystone/keystone/notifications.py:295}}
DEBUG keystone.notifications [-] Callback: `keystone.receipt.provider.Manager._drop_receipt_cache` subscribed to event `identity.invalidate_token_cache.internal`. {{(pid=222896) register_event_callback /opt/stack/keystone/keystone/notifications.py:295}}
DEBUG keystone.notifications [-] Callback: `keystone.trust.core.Manager._on_user_delete` subscribed to event `identity.user.deleted`. {{(pid=222896) register_event_callback /opt/stack/keystone/keystone/notifications.py:295}}
DEBUG keystone.notifications [-] Callback: `keystone.token.provider.Manager._drop_token_cache` subscribed to event `identity.OS-TRUST:trust.deleted`. {{(pid=222896) register_event_callback /opt/stack/keystone/keystone/notifications.py:295}}
DEBUG keystone.notifications [-] Callback: `keystone.token.provider.Manager._drop_token_cache` subscribed to event `identity.user.deleted`. {{(pid=222896) register_event_callback /opt/stack/keystone/keystone/notifications.py:295}}
DEBUG keystone.notifications [-] Callback: `keystone.token.provider.Manager._drop_token_cache` subscribed to event `identity.domain.deleted`. {{(pid=222896) register_event_callback /opt/stack/keystone/keystone/notifications.py:295}}
DEBUG keystone.notifications [-] Callback: `keystone.token.provider.Manager._drop_token_cache` subscribed to event `identity.user.disabled`. {{(pid=222896) register_event_callback /opt/stack/keystone/keystone/notifications.py:295}}
DEBUG keystone.notifications [-] Callback: `keystone.token.provider.Manager._drop_token_cache` subscribed to event `identity.domain.disabled`. {{(pid=222896) register_event_callback /opt/stack/keystone/keystone/notifications.py:295}}
DEBUG keystone.notifications [-] Callback: `keystone.token.provider.Manager._drop_token_cache` subscribed to event `identity.project.disabled`. {{(pid=222896) register_event_callback /opt/stack/keystone/keystone/notifications.py:295}}
DEBUG keystone.notifications [-] Callback: `keystone.token.provider.Manager._drop_token_cache` subscribed to event `identity.invalidate_token_cache.internal`. {{(pid=222896) register_event_callback /opt/stack/keystone/keystone/notifications.py:295}}
INFO dbcounter [-] Registered counter for database keystone
DEBUG oslo_db.sqlalchemy.engines [-] MySQL server mode set to STRICT_TRANS_TABLES,STRICT_ALL_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,TRADITIONAL,NO_ENGINE_SUBSTITUTION {{(pid=222896) _check_effective_sql_mode /opt/stack/data/venv/lib/python3.10/site-packages/oslo_db/sqlalchemy/engines.py:342}}
DEBUG dbcounter [-] [222896] Writer thread running {{(pid=222896) stat_writer /opt/stack/data/venv/lib/python3.10/site-packages/dbcounter.py:102}}
INFO keystone.cmd.bootstrap [None req-8376d365-5620-426e-9d6e-ccd97930ace4 None None] Created domain default
INFO keystone.cmd.bootstrap [None req-8376d365-5620-426e-9d6e-ccd97930ace4 None None] Created project admin
WARNING keystone.common.password_hashing [None req-8376d365-5620-426e-9d6e-ccd97930ace4 None None] Truncating password to algorithm specific maximum length 72 characters.: keystone.exception.UserNotFound: Could not find user: admin.
WARNING passlib.handlers.bcrypt [None req-8376d365-5620-426e-9d6e-ccd97930ace4 None None] (trapped) error reading bcrypt version: AttributeError: module 'bcrypt' has no attribute '__about__'
ERROR passlib.handlers.bcrypt Traceback (most recent call last):
ERROR passlib.handlers.bcrypt   File "/opt/stack/keystone/keystone/identity/backends/sql.py", line 209, in get_user_by_name
ERROR passlib.handlers.bcrypt     user_ref = query.one()
ERROR passlib.handlers.bcrypt   File "/opt/stack/data/venv/lib/python3.10/site-packages/sqlalchemy/orm/query.py", line 2870, in one
ERROR passlib.handlers.bcrypt     return self._iter().one()
ERROR passlib.handlers.bcrypt   File "/opt/stack/data/venv/lib/python3.10/site-packages/sqlalchemy/engine/result.py", line 1522, in one
ERROR passlib.handlers.bcrypt     return self._only_one_row(
ERROR passlib.handlers.bcrypt   File "/opt/stack/data/venv/lib/python3.10/site-packages/sqlalchemy/engine/result.py", line 562, in _only_one_row
ERROR passlib.handlers.bcrypt     raise exc.NoResultFound(
ERROR passlib.handlers.bcrypt sqlalchemy.exc.NoResultFound: No row was found when one was required
ERROR passlib.handlers.bcrypt 
ERROR passlib.handlers.bcrypt During handling of the above exception, another exception occurred:
ERROR passlib.handlers.bcrypt 
ERROR passlib.handlers.bcrypt Traceback (most recent call last):
ERROR passlib.handlers.bcrypt   File "/opt/stack/keystone/keystone/cmd/bootstrap.py", line 205, in _bootstrap_admin_user
ERROR passlib.handlers.bcrypt     user = PROVIDERS.identity_api.get_user_by_name(
ERROR passlib.handlers.bcrypt   File "/opt/stack/keystone/keystone/common/manager.py", line 110, in wrapped
ERROR passlib.handlers.bcrypt     __ret_val = __f(*args, **kwargs)
ERROR passlib.handlers.bcrypt   File "/opt/stack/keystone/keystone/identity/core.py", line 414, in wrapper
ERROR passlib.handlers.bcrypt     return f(self, *args, **kwargs)
ERROR passlib.handlers.bcrypt   File "/opt/stack/keystone/keystone/identity/core.py", line 424, in wrapper
ERROR passlib.handlers.bcrypt     return f(self, *args, **kwargs)
ERROR passlib.handlers.bcrypt   File "/opt/stack/data/venv/lib/python3.10/site-packages/decorator.py", line 232, in fun
ERROR passlib.handlers.bcrypt     return caller(func, *(extras + args), **kw)
ERROR passlib.handlers.bcrypt   File "/opt/stack/data/venv/lib/python3.10/site-packages/dogpile/cache/region.py", line 1632, in get_or_create_for_user_func
ERROR passlib.handlers.bcrypt     return self.get_or_create(
ERROR passlib.handlers.bcrypt   File "/opt/stack/data/venv/lib/python3.10/site-packages/dogpile/cache/region.py", line 1092, in get_or_create
ERROR passlib.handlers.bcrypt     with Lock(
ERROR passlib.handlers.bcrypt   File "/opt/stack/data/venv/lib/python3.10/site-packages/dogpile/lock.py", line 185, in __enter__
ERROR passlib.handlers.bcrypt     return self._enter()
ERROR passlib.handlers.bcrypt   File "/opt/stack/data/venv/lib/python3.10/site-packages/dogpile/lock.py", line 94, in _enter
ERROR passlib.handlers.bcrypt     generated = self._enter_create(value, createdtime)
ERROR passlib.handlers.bcrypt   File "/opt/stack/data/venv/lib/python3.10/site-packages/dogpile/lock.py", line 178, in _enter_create
ERROR passlib.handlers.bcrypt     return self.creator()
ERROR passlib.handlers.bcrypt   File "/opt/stack/data/venv/lib/python3.10/site-packages/dogpile/cache/region.py", line 1046, in gen_value
ERROR passlib.handlers.bcrypt     created_value = creator(
ERROR passlib.handlers.bcrypt   File "/opt/stack/keystone/keystone/identity/core.py", line 1035, in get_user_by_name
ERROR passlib.handlers.bcrypt     ref = driver.get_user_by_name(user_name, domain_id)
ERROR passlib.handlers.bcrypt   File "/opt/stack/keystone/keystone/identity/backends/sql.py", line 211, in get_user_by_name
ERROR passlib.handlers.bcrypt     raise exception.UserNotFound(user_id=user_name)
ERROR passlib.handlers.bcrypt keystone.exception.UserNotFound: Could not find user: admin.
ERROR passlib.handlers.bcrypt 
ERROR passlib.handlers.bcrypt During handling of the above exception, another exception occurred:
ERROR passlib.handlers.bcrypt 
ERROR passlib.handlers.bcrypt Traceback (most recent call last):
ERROR passlib.handlers.bcrypt   File "/opt/stack/data/venv/lib/python3.10/site-packages/passlib/handlers/bcrypt.py", line 620, in _load_backend_mixin
ERROR passlib.handlers.bcrypt     version = _bcrypt.__about__.__version__
ERROR passlib.handlers.bcrypt AttributeError: module 'bcrypt' has no attribute '__about__'
ERROR passlib.handlers.bcrypt 
DEBUG passlib.handlers.bcrypt [None req-8376d365-5620-426e-9d6e-ccd97930ace4 None None] detected 'bcrypt' backend, version '<unknown>' {{(pid=222896) _load_backend_mixin /opt/stack/data/venv/lib/python3.10/site-packages/passlib/handlers/bcrypt.py:625}}
DEBUG passlib.handlers.bcrypt [None req-8376d365-5620-426e-9d6e-ccd97930ace4 None None] 'bcrypt' backend lacks $2$ support, enabling workaround {{(pid=222896) _finalize_backend_mixin /opt/stack/data/venv/lib/python3.10/site-packages/passlib/handlers/bcrypt.py:406}}
INFO keystone.cmd.bootstrap [None req-8376d365-5620-426e-9d6e-ccd97930ace4 None None] Created user admin
INFO keystone.cmd.bootstrap [None req-8376d365-5620-426e-9d6e-ccd97930ace4 None None] Created role reader
INFO keystone.cmd.bootstrap [None req-8376d365-5620-426e-9d6e-ccd97930ace4 None None] Created role member
INFO keystone.cmd.bootstrap [None req-8376d365-5620-426e-9d6e-ccd97930ace4 None None] Created implied role where a0b0b31487d94ba09f896e4e97d6b605 implies 4c50c4544d7346a6bb31953109d431b9
INFO keystone.cmd.bootstrap [None req-8376d365-5620-426e-9d6e-ccd97930ace4 None None] Created role manager
INFO keystone.cmd.bootstrap [None req-8376d365-5620-426e-9d6e-ccd97930ace4 None None] Created implied role where f5009c0d0fe443ffb0e9c02a1a750f01 implies a0b0b31487d94ba09f896e4e97d6b605
INFO keystone.cmd.bootstrap [None req-8376d365-5620-426e-9d6e-ccd97930ace4 None None] Created role admin
INFO keystone.cmd.bootstrap [None req-8376d365-5620-426e-9d6e-ccd97930ace4 None None] Created implied role where 8ecab0167a544d758f0938f904d1b928 implies f5009c0d0fe443ffb0e9c02a1a750f01
INFO keystone.cmd.bootstrap [None req-8376d365-5620-426e-9d6e-ccd97930ace4 None None] Created role service
INFO keystone.cmd.bootstrap [None req-8376d365-5620-426e-9d6e-ccd97930ace4 None None] Granted role admin on project admin to user admin.
INFO keystone.cmd.bootstrap [None req-8376d365-5620-426e-9d6e-ccd97930ace4 None None] Granted role admin on the system to user admin.
WARNING py.warnings [None req-8376d365-5620-426e-9d6e-ccd97930ace4 None None] /opt/stack/data/venv/lib/python3.10/site-packages/pycadf/identifier.py:71: UserWarning: Invalid uuid: RegionOne. To ensure interoperability, identifiers should be a valid uuid.
  warnings.warn(('Invalid uuid: %s. To ensure interoperability, '

INFO keystone.cmd.bootstrap [None req-8376d365-5620-426e-9d6e-ccd97930ace4 None None] Created region RegionOne
INFO keystone.cmd.bootstrap [None req-8376d365-5620-426e-9d6e-ccd97930ace4 None None] Created public endpoint http://192.168.137.103/identity
+lib/keystone:bootstrap_keystone:615       '[' False == True ']'
+./stack.sh:main:1133                      create_keystone_accounts
+lib/keystone:create_keystone_accounts:314  local admin_project
++lib/keystone:create_keystone_accounts:315  oscwrap project show admin -f value -c id
++functions-common:oscwrap:2442             local xtrace
+++functions-common:oscwrap:2443             set +o
+++functions-common:oscwrap:2443             grep xtrace
++functions-common:oscwrap:2443             xtrace='set -o xtrace'
++functions-common:oscwrap:2444             set +o xtrace
++inc/async:async_inner:64                  rc=0
++inc/async:async_inner:65                  set +o xtrace
[217211 Async install_tempest:217211]: finished successfully
Failed to discover available identity versions when contacting http://192.168.137.103/identity. Attempting to parse version from URL.
Could not find versioned identity endpoints when attempting to authenticate. Please check that your auth_url is correct. Bad Gateway (HTTP 502)
++functions-common:oscwrap:2461             return 1
+lib/keystone:create_keystone_accounts:315  admin_project=
+lib/keystone:create_keystone_accounts:1   exit_trap
+./stack.sh:exit_trap:549                  local r=1
++./stack.sh:exit_trap:550                  jobs -p
+./stack.sh:exit_trap:550                  jobs=217211
+./stack.sh:exit_trap:553                  [[ -n 217211 ]]
+./stack.sh:exit_trap:553                  [[ -n '' ]]
+./stack.sh:exit_trap:559                  '[' -f /tmp/tmp.PfWC6Pj5hb ']'
+./stack.sh:exit_trap:560                  rm /tmp/tmp.PfWC6Pj5hb
+./stack.sh:exit_trap:564                  kill_spinner
+./stack.sh:kill_spinner:459               '[' '!' -z '' ']'
+./stack.sh:exit_trap:566                  [[ 1 -ne 0 ]]
+./stack.sh:exit_trap:567                  echo 'Error on exit'
Error on exit
+./stack.sh:exit_trap:569                  type -p generate-subunit
+./stack.sh:exit_trap:570                  generate-subunit 1714641451 1316 fail
+./stack.sh:exit_trap:572                  [[ -z /opt/stack/logs ]]
+./stack.sh:exit_trap:575                  /opt/stack/data/venv/bin/python3 /opt/stack/devstack/tools/worlddump.py -d /opt/stack/logs
World dumping... see /opt/stack/logs/worlddump-2024-05-02-093927.txt for details
+./stack.sh:exit_trap:584                  exit 1
stack@gpu-3:~/devstack$ . openrc 
WARNING: setting legacy OS_TENANT_NAME to support cli tools.
stack@gpu-3:~/devstack$ opestack server list
Command 'opestack' not found, did you mean:
  command 'openstack' from snap openstackclients (xena)
  command 'openstack' from deb python3-openstackclient (5.8.0-0ubuntu1)
See 'snap info <snapname>' for additional versions.
stack@gpu-3:~/devstack$ openstack server list
Failed to discover available identity versions when contacting http://192.168.137.103/identity. Attempting to parse version from URL.
Could not find versioned identity endpoints when attempting to authenticate. Please check that your auth_url is correct. Bad Gateway (HTTP 502)
stack@gpu-3:~/devstack$ 

