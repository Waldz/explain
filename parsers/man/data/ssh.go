package data

const SSH_OUTPUT = `.\"
.\" Author: Tatu Ylonen <ylo@cs.hut.fi>
.\" Copyright (c) 1995 Tatu Ylonen <ylo@cs.hut.fi>, Espoo, Finland
.\"                    All rights reserved
.\"
.\" As far as I am concerned, the code I have written for this software
.\" can be used freely for any purpose.  Any derived versions of this
.\" software must be clearly marked as such, and if the derived work is
.\" incompatible with the protocol description in the RFC file, it must be
.\" called by a name other than "ssh" or "Secure Shell".
.\"
.\" Copyright (c) 1999,2000 Markus Friedl.  All rights reserved.
.\" Copyright (c) 1999 Aaron Campbell.  All rights reserved.
.\" Copyright (c) 1999 Theo de Raadt.  All rights reserved.
.\"
.\" Redistribution and use in source and binary forms, with or without
.\" modification, are permitted provided that the following conditions
.\" are met:
.\" 1. Redistributions of source code must retain the above copyright
.\"    notice, this list of conditions and the following disclaimer.
.\" 2. Redistributions in binary form must reproduce the above copyright
.\"    notice, this list of conditions and the following disclaimer in the
.\"    documentation and/or other materials provided with the distribution.
.\"
.\" THIS SOFTWARE IS PROVIDED BY THE AUTHOR AS IS'' AND ANY EXPRESS OR
.\" IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES
.\" OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED.
.\" IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR ANY DIRECT, INDIRECT,
.\" INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT
.\" NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
.\" DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
.\" THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
.\" (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF
.\" THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
.\"
.\" $OpenBSD: ssh.1,v 1.403 2019/06/12 11:31:50 jmc Exp $
.Dd $Mdocdate: June 12 2019 $
.Dt SSH 1
.Os
.Sh NAME
.Nm ssh
.Nd OpenSSH SSH client (remote login program)
.Sh SYNOPSIS
.Nm ssh
.Op Fl 46AaCfGgKkMNnqsTtVvXxYy
.Op Fl B Ar bind_interface
.Op Fl b Ar bind_address
.Op Fl c Ar cipher_spec
.Op Fl D Oo Ar bind_address : Oc Ns Ar port
.Op Fl E Ar log_file
.Op Fl e Ar escape_char
.Op Fl F Ar configfile
.Op Fl I Ar pkcs11
.Op Fl i Ar identity_file
.Op Fl J Ar destination
.Op Fl L Ar address
.Op Fl l Ar login_name
.Op Fl m Ar mac_spec
.Op Fl O Ar ctl_cmd
.Op Fl o Ar option
.Op Fl p Ar port
.Op Fl Q Ar query_option
.Op Fl R Ar address
.Op Fl S Ar ctl_path
.Op Fl W Ar host : Ns Ar port
.Op Fl w Ar local_tun Ns Op : Ns Ar remote_tun
.Ar destination
.Op Ar command
.Sh DESCRIPTION
.Nm
(SSH client) is a program for logging into a remote machine and for
executing commands on a remote machine.
It is intended to provide secure encrypted communications between
two untrusted hosts over an insecure network.
X11 connections, arbitrary TCP ports and
.Ux Ns -domain
sockets can also be forwarded over the secure channel.
.Pp
.Nm
connects and logs into the specified
.Ar destination ,
which may be specified as either
.Sm off
.Oo user @ Oc hostname
.Sm on
or a URI of the form
.Sm off
.No ssh:// Oo user @ Oc hostname Op : port .
.Sm on
The user must prove
his/her identity to the remote machine using one of several methods
(see below).
.Pp
If a
.Ar command
is specified,
it is executed on the remote host instead of a login shell.
.Pp
The options are as follows:
.Pp
.Bl -tag -width Ds -compact
.It Fl 4
Forces
.Nm
to use IPv4 addresses only.
.Pp
.It Fl 6
Forces
.Nm
to use IPv6 addresses only.
.Pp
.It Fl A
Enables forwarding of the authentication agent connection.
This can also be specified on a per-host basis in a configuration file.
.Pp
Agent forwarding should be enabled with caution.
Users with the ability to bypass file permissions on the remote host
(for the agent's
.Ux Ns -domain
socket) can access the local agent through the forwarded connection.
An attacker cannot obtain key material from the agent,
however they can perform operations on the keys that enable them to
authenticate using the identities loaded into the agent.
.Pp
.It Fl a
Disables forwarding of the authentication agent connection.
.Pp
.It Fl B Ar bind_interface
Bind to the address of
.Ar bind_interface
before attempting to connect to the destination host.
This is only useful on systems with more than one address.
.Pp
.It Fl b Ar bind_address
Use
.Ar bind_address
on the local machine as the source address
of the connection.
Only useful on systems with more than one address.
.Pp
.It Fl C
Requests compression of all data (including stdin, stdout, stderr, and
data for forwarded X11, TCP and
.Ux Ns -domain
connections).
The compression algorithm is the same used by
.Xr gzip 1 .
Compression is desirable on modem lines and other
slow connections, but will only slow down things on fast networks.
The default value can be set on a host-by-host basis in the
configuration files; see the
.Cm Compression
option.
.Pp
.It Fl c Ar cipher_spec
Selects the cipher specification for encrypting the session.
.Ar cipher_spec
is a comma-separated list of ciphers
listed in order of preference.
See the
.Cm Ciphers
keyword in
.Xr ssh_config 5
for more information.
.Pp
.It Fl D Xo
.Sm off
.Oo Ar bind_address : Oc
.Ar port
.Sm on
.Xc
Specifies a local
.Dq dynamic
application-level port forwarding.
This works by allocating a socket to listen to
.Ar port
on the local side, optionally bound to the specified
.Ar bind_address .
Whenever a connection is made to this port, the
connection is forwarded over the secure channel, and the application
protocol is then used to determine where to connect to from the
remote machine.
Currently the SOCKS4 and SOCKS5 protocols are supported, and
.Nm
will act as a SOCKS server.
Only root can forward privileged ports.
Dynamic port forwardings can also be specified in the configuration file.
.Pp
IPv6 addresses can be specified by enclosing the address in square brackets.
Only the superuser can forward privileged ports.
By default, the local port is bound in accordance with the
.Cm GatewayPorts
setting.
However, an explicit
.Ar bind_address
may be used to bind the connection to a specific address.
The
.Ar bind_address
of
.Dq localhost
indicates that the listening port be bound for local use only, while an
empty address or
.Sq *
indicates that the port should be available from all interfaces.
.Pp
.It Fl E Ar log_file
Append debug logs to
.Ar log_file
instead of standard error.
.Pp
.It Fl e Ar escape_char
Sets the escape character for sessions with a pty (default:
.Ql ~ ) .
The escape character is only recognized at the beginning of a line.
The escape character followed by a dot
.Pq Ql \&.
closes the connection;
followed by control-Z suspends the connection;
and followed by itself sends the escape character once.
Setting the character to
.Dq none
disables any escapes and makes the session fully transparent.
.Pp
.It Fl F Ar configfile
Specifies an alternative per-user configuration file.
If a configuration file is given on the command line,
the system-wide configuration file
.Pq Pa /etc/ssh/ssh_config
will be ignored.
The default for the per-user configuration file is
.Pa ~/.ssh/config .
.Pp
.It Fl f
Requests
.Nm
to go to background just before command execution.
This is useful if
.Nm
is going to ask for passwords or passphrases, but the user
wants it in the background.
This implies
.Fl n .
The recommended way to start X11 programs at a remote site is with
something like
.Ic ssh -f host xterm .
.Pp
If the
.Cm ExitOnForwardFailure
configuration option is set to
.Dq yes ,
then a client started with
.Fl f
will wait for all remote port forwards to be successfully established
before placing itself in the background.
.Pp
.It Fl G
Causes
.Nm
to print its configuration after evaluating
.Cm Host
and
.Cm Match
blocks and exit.
.Pp
.It Fl g
Allows remote hosts to connect to local forwarded ports.
If used on a multiplexed connection, then this option must be specified
on the master process.
.Pp
.It Fl I Ar pkcs11
Specify the PKCS#11 shared library
.Nm
should use to communicate with a PKCS#11 token providing keys for user
authentication.
.Pp
.It Fl i Ar identity_file
Selects a file from which the identity (private key) for
public key authentication is read.
The default is
.Pa ~/.ssh/id_dsa ,
.Pa ~/.ssh/id_ecdsa ,
.Pa ~/.ssh/id_ed25519
and
.Pa ~/.ssh/id_rsa .
Identity files may also be specified on
a per-host basis in the configuration file.
It is possible to have multiple
.Fl i
options (and multiple identities specified in
configuration files).
If no certificates have been explicitly specified by the
.Cm CertificateFile
directive,
.Nm
will also try to load certificate information from the filename obtained
by appending
.Pa -cert.pub
to identity filenames.
.Pp
.It Fl J Ar destination
Connect to the target host by first making a
.Nm
connection to the jump host described by
.Ar destination
and then establishing a TCP forwarding to the ultimate destination from
there.
Multiple jump hops may be specified separated by comma characters.
This is a shortcut to specify a
.Cm ProxyJump
configuration directive.
Note that configuration directives supplied on the command-line generally
apply to the destination host and not any specified jump hosts.
Use
.Pa ~/.ssh/config
to specify configuration for jump hosts.
.Pp
.It Fl K
Enables GSSAPI-based authentication and forwarding (delegation) of GSSAPI
credentials to the server.
.Pp
.It Fl k
Disables forwarding (delegation) of GSSAPI credentials to the server.
.Pp
.It Fl L Xo
.Sm off
.Oo Ar bind_address : Oc
.Ar port : host : hostport
.Sm on
.Xc
.It Fl L Xo
.Sm off
.Oo Ar bind_address : Oc
.Ar port : remote_socket
.Sm on
.Xc
.It Fl L Xo
.Sm off
.Ar local_socket : host : hostport
.Sm on
.Xc
.It Fl L Xo
.Sm off
.Ar local_socket : remote_socket
.Sm on
.Xc
Specifies that connections to the given TCP port or Unix socket on the local
(client) host are to be forwarded to the given host and port, or Unix socket,
on the remote side.
This works by allocating a socket to listen to either a TCP
.Ar port
on the local side, optionally bound to the specified
.Ar bind_address ,
or to a Unix socket.
Whenever a connection is made to the local port or socket, the
connection is forwarded over the secure channel, and a connection is
made to either
.Ar host
port
.Ar hostport ,
or the Unix socket
.Ar remote_socket ,
from the remote machine.
.Pp
Port forwardings can also be specified in the configuration file.
Only the superuser can forward privileged ports.
IPv6 addresses can be specified by enclosing the address in square brackets.
.Pp
By default, the local port is bound in accordance with the
.Cm GatewayPorts
setting.
However, an explicit
.Ar bind_address
may be used to bind the connection to a specific address.
The
.Ar bind_address
of
.Dq localhost
indicates that the listening port be bound for local use only, while an
empty address or
.Sq *
indicates that the port should be available from all interfaces.
.Pp
.It Fl l Ar login_name
Specifies the user to log in as on the remote machine.
This also may be specified on a per-host basis in the configuration file.
.Pp
.It Fl M
Places the
.Nm
client into
.Dq master
mode for connection sharing.
Multiple
.Fl M
options places
.Nm
into
.Dq master
mode but with confirmation required using
.Xr ssh-askpass 1
before each operation that changes the multiplexing state
(e.g. opening a new session).
Refer to the description of
.Cm ControlMaster
in
.Xr ssh_config 5
for details.
.Pp
.It Fl m Ar mac_spec
A comma-separated list of MAC (message authentication code) algorithms,
specified in order of preference.
See the
.Cm MACs
keyword for more information.
.Pp
.It Fl N
Do not execute a remote command.
This is useful for just forwarding ports.
.Pp
.It Fl n
Redirects stdin from
.Pa /dev/null
(actually, prevents reading from stdin).
This must be used when
.Nm
is run in the background.
A common trick is to use this to run X11 programs on a remote machine.
For example,
.Ic ssh -n shadows.cs.hut.fi emacs &
will start an emacs on shadows.cs.hut.fi, and the X11
connection will be automatically forwarded over an encrypted channel.
The
.Nm
program will be put in the background.
(This does not work if
.Nm
needs to ask for a password or passphrase; see also the
.Fl f
option.)
.Pp
.It Fl O Ar ctl_cmd
Control an active connection multiplexing master process.
When the
.Fl O
option is specified, the
.Ar ctl_cmd
argument is interpreted and passed to the master process.
Valid commands are:
.Dq check
(check that the master process is running),
.Dq forward
(request forwardings without command execution),
.Dq cancel
(cancel forwardings),
.Dq exit
(request the master to exit), and
.Dq stop
(request the master to stop accepting further multiplexing requests).
.Pp
.It Fl o Ar option
Can be used to give options in the format used in the configuration file.
This is useful for specifying options for which there is no separate
command-line flag.
For full details of the options listed below, and their possible values, see
.Xr ssh_config 5 .
.Pp
.Bl -tag -width Ds -offset indent -compact
.It AddKeysToAgent
.It AddressFamily
.It BatchMode
.It BindAddress
.It CanonicalDomains
.It CanonicalizeFallbackLocal
.It CanonicalizeHostname
.It CanonicalizeMaxDots
.It CanonicalizePermittedCNAMEs
.It CASignatureAlgorithms
.It CertificateFile
.It ChallengeResponseAuthentication
.It CheckHostIP
.It Ciphers
.It ClearAllForwardings
.It Compression
.It ConnectionAttempts
.It ConnectTimeout
.It ControlMaster
.It ControlPath
.It ControlPersist
.It DynamicForward
.It EscapeChar
.It ExitOnForwardFailure
.It FingerprintHash
.It ForwardAgent
.It ForwardX11
.It ForwardX11Timeout
.It ForwardX11Trusted
.It GatewayPorts
.It GlobalKnownHostsFile
.It GSSAPIAuthentication
.It GSSAPIDelegateCredentials
.It HashKnownHosts
.It Host
.It HostbasedAuthentication
.It HostbasedKeyTypes
.It HostKeyAlgorithms
.It HostKeyAlias
.It Hostname
.It IdentitiesOnly
.It IdentityAgent
.It IdentityFile
.It IPQoS
.It KbdInteractiveAuthentication
.It KbdInteractiveDevices
.It KexAlgorithms
.It LocalCommand
.It LocalForward
.It LogLevel
.It MACs
.It Match
.It NoHostAuthenticationForLocalhost
.It NumberOfPasswordPrompts
.It PasswordAuthentication
.It PermitLocalCommand
.It PKCS11Provider
.It Port
.It PreferredAuthentications
.It ProxyCommand
.It ProxyJump
.It ProxyUseFdpass
.It PubkeyAcceptedKeyTypes
.It PubkeyAuthentication
.It RekeyLimit
.It RemoteCommand
.It RemoteForward
.It RequestTTY
.It SendEnv
.It ServerAliveInterval
.It ServerAliveCountMax
.It SetEnv
.It StreamLocalBindMask
.It StreamLocalBindUnlink
.It StrictHostKeyChecking
.It TCPKeepAlive
.It Tunnel
.It TunnelDevice
.It UpdateHostKeys
.\" #ifdef __APPLE_KEYCHAIN__
.It UseKeychain
.\" #endif
.It User
.It UserKnownHostsFile
.It VerifyHostKeyDNS
.It VisualHostKey
.It XAuthLocation
.El
.Pp
.It Fl p Ar port
Port to connect to on the remote host.
This can be specified on a
per-host basis in the configuration file.
.Pp
.It Fl Q Ar query_option
Queries
.Nm
for the algorithms supported for the specified version 2.
The available features are:
.Ar cipher
(supported symmetric ciphers),
.Ar cipher-auth
(supported symmetric ciphers that support authenticated encryption),
.Ar help
(supported query terms for use with the
.Fl Q
flag),
.Ar mac
(supported message integrity codes),
.Ar kex
(key exchange algorithms),
.Ar key
(key types),
.Ar key-cert
(certificate key types),
.Ar key-plain
(non-certificate key types),
.Ar protocol-version
(supported SSH protocol versions), and
.Ar sig
(supported signature algorithms).
.Pp
.It Fl q
Quiet mode.
Causes most warning and diagnostic messages to be suppressed.
.Pp
.It Fl R Xo
.Sm off
.Oo Ar bind_address : Oc
.Ar port : host : hostport
.Sm on
.Xc
[bind_address:]port:host:hostport
.It Fl R Xo
.Sm off
.Oo Ar bind_address : Oc
.Ar port : local_socket
.Sm on
.Xc
.It Fl R Xo
.Sm off
.Ar remote_socket : host : hostport
.Sm on
.Xc
.It Fl R Xo
.Sm off
.Ar remote_socket : local_socket
.Sm on
.Xc
.It Fl R Xo
.Sm off
.Oo Ar bind_address : Oc
.Ar port
.Sm on
.Xc
[bind_address:]port:host:hostport
Specifies that connections to the given TCP port or Unix socket on the remote
(server) host are to be forwarded to the local side.
.Pp
This works by allocating a socket to listen to either a TCP
.Ar port
or to a Unix socket on the remote side.
Whenever a connection is made to this port or Unix socket, the
connection is forwarded over the secure channel, and a connection
is made from the local machine to either an explicit destination specified by
.Ar host
port
.Ar hostport ,
or
.Ar local_socket ,
or, if no explicit destination was specified,
.Nm
will act as a SOCKS 4/5 proxy and forward connections to the destinations
requested by the remote SOCKS client.
.Pp
Port forwardings can also be specified in the configuration file.
Privileged ports can be forwarded only when
logging in as root on the remote machine.
IPv6 addresses can be specified by enclosing the address in square brackets.
.Pp
By default, TCP listening sockets on the server will be bound to the loopback
interface only.
This may be overridden by specifying a
.Ar bind_address .
An empty
.Ar bind_address ,
or the address
.Ql * ,
indicates that the remote socket should listen on all interfaces.
Specifying a remote
.Ar bind_address
will only succeed if the server's
.Cm GatewayPorts
option is enabled (see
.Xr sshd_config 5 ) .
.Pp
If the
.Ar port
argument is
.Ql 0 ,
the listen port will be dynamically allocated on the server and reported
to the client at run time.
When used together with
.Ic -O forward
the allocated port will be printed to the standard output.
.Pp
.It Fl S Ar ctl_path
Specifies the location of a control socket for connection sharing,
or the string
.Dq none
to disable connection sharing.
Refer to the description of
.Cm ControlPath
and
.Cm ControlMaster
in
.Xr ssh_config 5
for details.
.Pp
.It Fl s
May be used to request invocation of a subsystem on the remote system.
Subsystems facilitate the use of SSH
as a secure transport for other applications (e.g.\&
.Xr sftp 1 ) .
The subsystem is specified as the remote command.
.Pp
.It Fl T
Disable pseudo-terminal allocation.
.Pp
.It Fl t
Force pseudo-terminal allocation.
This can be used to execute arbitrary
screen-based programs on a remote machine, which can be very useful,
e.g. when implementing menu services.
Multiple
.Fl t
options force tty allocation, even if
.Nm
has no local tty.
.Pp
.It Fl V
Display the version number and exit.
.Pp
.It Fl v
Verbose mode.
Causes
.Nm
to print debugging messages about its progress.
This is helpful in
debugging connection, authentication, and configuration problems.
Multiple
.Fl v
options increase the verbosity.
The maximum is 3.
.Pp
.It Fl W Ar host : Ns Ar port
Requests that standard input and output on the client be forwarded to
.Ar host
on
.Ar port
over the secure channel.
Implies
.Fl N ,
.Fl T ,
.Cm ExitOnForwardFailure
and
.Cm ClearAllForwardings ,
though these can be overridden in the configuration file or using
.Fl o
command line options.
.Pp
.It Fl w Xo
.Ar local_tun Ns Op : Ns Ar remote_tun
.Xc
Requests
tunnel
device forwarding with the specified
.Xr tun 4
devices between the client
.Pq Ar local_tun
and the server
.Pq Ar remote_tun .
.Pp
The devices may be specified by numerical ID or the keyword
.Dq any ,
which uses the next available tunnel device.
If
.Ar remote_tun
is not specified, it defaults to
.Dq any .
See also the
.Cm Tunnel
and
.Cm TunnelDevice
directives in
.Xr ssh_config 5 .
.Pp
If the
.Cm Tunnel
directive is unset, it will be set to the default tunnel mode, which is
.Dq point-to-point .
If a different
.Cm Tunnel
forwarding mode it desired, then it should be specified before
.Fl w .
.Pp
.It Fl X
Enables X11 forwarding.
This can also be specified on a per-host basis in a configuration file.
.Pp
X11 forwarding should be enabled with caution.
Users with the ability to bypass file permissions on the remote host
(for the user's X authorization database)
can access the local X11 display through the forwarded connection.
An attacker may then be able to perform activities such as keystroke monitoring.
.Pp
For this reason, X11 forwarding is subjected to X11 SECURITY extension
restrictions by default.
Please refer to the
.Nm
.Fl Y
option and the
.Cm ForwardX11Trusted
directive in
.Xr ssh_config 5
for more information.
.Pp
.It Fl x
Disables X11 forwarding.
.Pp
.It Fl Y
Enables trusted X11 forwarding.
Trusted X11 forwardings are not subjected to the X11 SECURITY extension
controls.
.Pp
.It Fl y
Send log information using the
.Xr syslog 3
system module.
By default this information is sent to stderr.
.El
.Pp
.Nm
may additionally obtain configuration data from
a per-user configuration file and a system-wide configuration file.
The file format and configuration options are described in
.Xr ssh_config 5 .
.Sh AUTHENTICATION
The OpenSSH SSH client supports SSH protocol 2.
.Pp
The methods available for authentication are:
GSSAPI-based authentication,
host-based authentication,
public key authentication,
challenge-response authentication,
and password authentication.
Authentication methods are tried in the order specified above,
though
.Cm PreferredAuthentications
can be used to change the default order.
.Pp
Host-based authentication works as follows:
If the machine the user logs in from is listed in
.Pa /etc/hosts.equiv
or
.Pa /etc/shosts.equiv
on the remote machine, and the user names are
the same on both sides, or if the files
.Pa ~/.rhosts
or
.Pa ~/.shosts
exist in the user's home directory on the
remote machine and contain a line containing the name of the client
machine and the name of the user on that machine, the user is
considered for login.
Additionally, the server
.Em must
be able to verify the client's
host key (see the description of
.Pa /etc/ssh/ssh_known_hosts
and
.Pa ~/.ssh/known_hosts ,
below)
for login to be permitted.
This authentication method closes security holes due to IP
spoofing, DNS spoofing, and routing spoofing.
[Note to the administrator:
.Pa /etc/hosts.equiv ,
.Pa ~/.rhosts ,
and the rlogin/rsh protocol in general, are inherently insecure and should be
disabled if security is desired.]
.Pp
Public key authentication works as follows:
The scheme is based on public-key cryptography,
using cryptosystems
where encryption and decryption are done using separate keys,
and it is unfeasible to derive the decryption key from the encryption key.
The idea is that each user creates a public/private
key pair for authentication purposes.
The server knows the public key, and only the user knows the private key.
.Nm
implements public key authentication protocol automatically,
using one of the DSA, ECDSA, Ed25519 or RSA algorithms.
The HISTORY section of
.Xr ssl 8
contains a brief discussion of the DSA and RSA algorithms.
.Pp
The file
.Pa ~/.ssh/authorized_keys
lists the public keys that are permitted for logging in.
When the user logs in, the
.Nm
program tells the server which key pair it would like to use for
authentication.
The client proves that it has access to the private key
and the server checks that the corresponding public key
is authorized to accept the account.
.Pp
The server may inform the client of errors that prevented public key
authentication from succeeding after authentication completes using a
different method.
These may be viewed by increasing the
.Cm LogLevel
to
.Cm DEBUG
or higher (e.g. by using the
.Fl v
flag).
.Pp
The user creates his/her key pair by running
.Xr ssh-keygen 1 .
This stores the private key in
.Pa ~/.ssh/id_dsa
(DSA),
.Pa ~/.ssh/id_ecdsa
(ECDSA),
.Pa ~/.ssh/id_ed25519
(Ed25519),
or
.Pa ~/.ssh/id_rsa
(RSA)
and stores the public key in
.Pa ~/.ssh/id_dsa.pub
(DSA),
.Pa ~/.ssh/id_ecdsa.pub
(ECDSA),
.Pa ~/.ssh/id_ed25519.pub
(Ed25519),
or
.Pa ~/.ssh/id_rsa.pub
(RSA)
in the user's home directory.
The user should then copy the public key
to
.Pa ~/.ssh/authorized_keys
in his/her home directory on the remote machine.
The
.Pa authorized_keys
file corresponds to the conventional
.Pa ~/.rhosts
file, and has one key
per line, though the lines can be very long.
After this, the user can log in without giving the password.
.Pp
A variation on public key authentication
is available in the form of certificate authentication:
instead of a set of public/private keys,
signed certificates are used.
This has the advantage that a single trusted certification authority
can be used in place of many public/private keys.
See the CERTIFICATES section of
.Xr ssh-keygen 1
for more information.
.Pp
The most convenient way to use public key or certificate authentication
may be with an authentication agent.
See
.Xr ssh-agent 1
and (optionally) the
.Cm AddKeysToAgent
directive in
.Xr ssh_config 5
for more information.
.Pp
Challenge-response authentication works as follows:
The server sends an arbitrary
.Qq challenge
text, and prompts for a response.
Examples of challenge-response authentication include
.Bx
Authentication (see
.Xr login.conf 5 )
and PAM (some
.Pf non- Ox
systems).
.Pp
Finally, if other authentication methods fail,
.Nm
prompts the user for a password.
The password is sent to the remote
host for checking; however, since all communications are encrypted,
the password cannot be seen by someone listening on the network.
.Pp
.Nm
automatically maintains and checks a database containing
identification for all hosts it has ever been used with.
Host keys are stored in
.Pa ~/.ssh/known_hosts
in the user's home directory.
Additionally, the file
.Pa /etc/ssh/ssh_known_hosts
is automatically checked for known hosts.
Any new hosts are automatically added to the user's file.
If a host's identification ever changes,
.Nm
warns about this and disables password authentication to prevent
server spoofing or man-in-the-middle attacks,
which could otherwise be used to circumvent the encryption.
The
.Cm StrictHostKeyChecking
option can be used to control logins to machines whose
host key is not known or has changed.
.Pp
When the user's identity has been accepted by the server, the server
either executes the given command in a non-interactive session or,
if no command has been specified, logs into the machine and gives
the user a normal shell as an interactive session.
All communication with
the remote command or shell will be automatically encrypted.
.Pp
If an interactive session is requested
.Nm
by default will only request a pseudo-terminal (pty) for interactive
sessions when the client has one.
The flags
.Fl T
and
.Fl t
can be used to override this behaviour.
.Pp
If a pseudo-terminal has been allocated the
user may use the escape characters noted below.
.Pp
If no pseudo-terminal has been allocated,
the session is transparent and can be used to reliably transfer binary data.
On most systems, setting the escape character to
.Dq none
will also make the session transparent even if a tty is used.
.Pp
The session terminates when the command or shell on the remote
machine exits and all X11 and TCP connections have been closed.
.Sh ESCAPE CHARACTERS
When a pseudo-terminal has been requested,
.Nm
supports a number of functions through the use of an escape character.
.Pp
A single tilde character can be sent as
.Ic ~~
or by following the tilde by a character other than those described below.
The escape character must always follow a newline to be interpreted as
special.
The escape character can be changed in configuration files using the
.Cm EscapeChar
configuration directive or on the command line by the
.Fl e
option.
.Pp
The supported escapes (assuming the default
.Ql ~ )
are:
.Bl -tag -width Ds
.It Cm ~.
Disconnect.
.It Cm ~^Z
Background
.Nm .
.It Cm ~#
List forwarded connections.
.It Cm ~&
Background
.Nm
at logout when waiting for forwarded connection / X11 sessions to terminate.
.It Cm ~?
Display a list of escape characters.
.It Cm ~B
Send a BREAK to the remote system
(only useful if the peer supports it).
.It Cm ~C
Open command line.
Currently this allows the addition of port forwardings using the
.Fl L ,
.Fl R
and
.Fl D
options (see above).
It also allows the cancellation of existing port-forwardings
with
.Sm off
.Fl KL Oo Ar bind_address : Oc Ar port
.Sm on
for local,
.Sm off
.Fl KR Oo Ar bind_address : Oc Ar port
.Sm on
for remote and
.Sm off
.Fl KD Oo Ar bind_address : Oc Ar port
.Sm on
for dynamic port-forwardings.
.Ic !\& Ns Ar command
allows the user to execute a local command if the
.Ic PermitLocalCommand
option is enabled in
.Xr ssh_config 5 .
Basic help is available, using the
.Fl h
option.
.It Cm ~R
Request rekeying of the connection
(only useful if the peer supports it).
.It Cm ~V
Decrease the verbosity
.Pq Ic LogLevel
when errors are being written to stderr.
.It Cm ~v
Increase the verbosity
.Pq Ic LogLevel
when errors are being written to stderr.
.El
.Sh TCP FORWARDING
Forwarding of arbitrary TCP connections over a secure channel
can be specified either on the command line or in a configuration file.
One possible application of TCP forwarding is a secure connection to a
mail server; another is going through firewalls.
.Pp
In the example below, we look at encrypting communication for an IRC client,
even though the IRC server it connects to does not directly
support encrypted communication.
This works as follows:
the user connects to the remote host using
.Nm ,
specifying the ports to be used to forward the connection.
After that it is possible to start the program locally,
and
.Nm
will encrypt and forward the connection to the remote server.
.Pp
The following example tunnels an IRC session from the client
to an IRC server at
.Dq server.example.com ,
joining channel
.Dq #users ,
nickname
.Dq pinky ,
using the standard IRC port, 6667:
.Bd -literal -offset 4n
$ ssh -f -L 6667:localhost:6667 server.example.com sleep 10
$ irc -c '#users' pinky IRC/127.0.0.1
.Ed
.Pp
The
.Fl f
option backgrounds
.Nm
and the remote command
.Dq sleep 10
is specified to allow an amount of time
(10 seconds, in the example)
to start the program which is going to use the tunnel.
If no connections are made within the time specified,
.Nm
will exit.
.Sh X11 FORWARDING
If the
.Cm ForwardX11
variable is set to
.Dq yes
(or see the description of the
.Fl X ,
.Fl x ,
and
.Fl Y
options above)
and the user is using X11 (the
.Ev DISPLAY
environment variable is set), the connection to the X11 display is
automatically forwarded to the remote side in such a way that any X11
programs started from the shell (or command) will go through the
encrypted channel, and the connection to the real X server will be made
from the local machine.
The user should not manually set
.Ev DISPLAY .
Forwarding of X11 connections can be
configured on the command line or in configuration files.
.Pp
The
.Ev DISPLAY
value set by
.Nm
will point to the server machine, but with a display number greater than zero.
This is normal, and happens because
.Nm
creates a
.Dq proxy
X server on the server machine for forwarding the
connections over the encrypted channel.
.Pp
.Nm
will also automatically set up Xauthority data on the server machine.
For this purpose, it will generate a random authorization cookie,
store it in Xauthority on the server, and verify that any forwarded
connections carry this cookie and replace it by the real cookie when
the connection is opened.
The real authentication cookie is never
sent to the server machine (and no cookies are sent in the plain).
.Pp
If the
.Cm ForwardAgent
variable is set to
.Dq yes
(or see the description of the
.Fl A
and
.Fl a
options above) and
the user is using an authentication agent, the connection to the agent
is automatically forwarded to the remote side.
.Sh VERIFYING HOST KEYS
When connecting to a server for the first time,
a fingerprint of the server's public key is presented to the user
(unless the option
.Cm StrictHostKeyChecking
has been disabled).
Fingerprints can be determined using
.Xr ssh-keygen 1 :
.Pp
.Dl $ ssh-keygen -l -f /etc/ssh/ssh_host_rsa_key
.Pp
If the fingerprint is already known, it can be matched
and the key can be accepted or rejected.
If only legacy (MD5) fingerprints for the server are available, the
.Xr ssh-keygen 1
.Fl E
option may be used to downgrade the fingerprint algorithm to match.
.Pp
Because of the difficulty of comparing host keys
just by looking at fingerprint strings,
there is also support to compare host keys visually,
using
.Em random art .
By setting the
.Cm VisualHostKey
option to
.Dq yes ,
a small ASCII graphic gets displayed on every login to a server, no matter
if the session itself is interactive or not.
By learning the pattern a known server produces, a user can easily
find out that the host key has changed when a completely different pattern
is displayed.
Because these patterns are not unambiguous however, a pattern that looks
similar to the pattern remembered only gives a good probability that the
host key is the same, not guaranteed proof.
.Pp
To get a listing of the fingerprints along with their random art for
all known hosts, the following command line can be used:
.Pp
.Dl $ ssh-keygen -lv -f ~/.ssh/known_hosts
.Pp
If the fingerprint is unknown,
an alternative method of verification is available:
SSH fingerprints verified by DNS.
An additional resource record (RR),
SSHFP,
is added to a zonefile
and the connecting client is able to match the fingerprint
with that of the key presented.
.Pp
In this example, we are connecting a client to a server,
.Dq host.example.com .
The SSHFP resource records should first be added to the zonefile for
host.example.com:
.Bd -literal -offset indent
$ ssh-keygen -r host.example.com.
.Ed
.Pp
The output lines will have to be added to the zonefile.
To check that the zone is answering fingerprint queries:
.Pp
.Dl $ dig -t SSHFP host.example.com
.Pp
Finally the client connects:
.Bd -literal -offset indent
$ ssh -o "VerifyHostKeyDNS ask" host.example.com
[...]
Matching host key fingerprint found in DNS.
Are you sure you want to continue connecting (yes/no)?
.Ed
.Pp
See the
.Cm VerifyHostKeyDNS
option in
.Xr ssh_config 5
for more information.
.Sh SSH-BASED VIRTUAL PRIVATE NETWORKS
.Nm
contains support for Virtual Private Network (VPN) tunnelling
using the
.Xr tun 4
network pseudo-device,
allowing two networks to be joined securely.
The
.Xr sshd_config 5
configuration option
.Cm PermitTunnel
controls whether the server supports this,
and at what level (layer 2 or 3 traffic).
.Pp
The following example would connect client network 10.0.50.0/24
with remote network 10.0.99.0/24 using a point-to-point connection
from 10.1.1.1 to 10.1.1.2,
provided that the SSH server running on the gateway to the remote network,
at 192.168.1.15, allows it.
.Pp
On the client:
.Bd -literal -offset indent
# ssh -f -w 0:1 192.168.1.15 true
# ifconfig tun0 10.1.1.1 10.1.1.2 netmask 255.255.255.252
# route add 10.0.99.0/24 10.1.1.2
.Ed
.Pp
On the server:
.Bd -literal -offset indent
# ifconfig tun1 10.1.1.2 10.1.1.1 netmask 255.255.255.252
# route add 10.0.50.0/24 10.1.1.1
.Ed
.Pp
Client access may be more finely tuned via the
.Pa /root/.ssh/authorized_keys
file (see below) and the
.Cm PermitRootLogin
server option.
The following entry would permit connections on
.Xr tun 4
device 1 from user
.Dq jane
and on tun device 2 from user
.Dq john ,
if
.Cm PermitRootLogin
is set to
.Dq forced-commands-only :
.Bd -literal -offset 2n
tunnel="1",command="sh /etc/netstart tun1" ssh-rsa ... jane
tunnel="2",command="sh /etc/netstart tun2" ssh-rsa ... john
.Ed
.Pp
Since an SSH-based setup entails a fair amount of overhead,
it may be more suited to temporary setups,
such as for wireless VPNs.
More permanent VPNs are better provided by tools such as
.Xr ipsecctl 8
and
.Xr isakmpd 8 .
.Sh ENVIRONMENT
.Nm
will normally set the following environment variables:
.Bl -tag -width "SSH_ORIGINAL_COMMAND"
.It Ev DISPLAY
The
.Ev DISPLAY
variable indicates the location of the X11 server.
It is automatically set by
.Nm
to point to a value of the form
.Dq hostname:n ,
where
.Dq hostname
indicates the host where the shell runs, and
.Sq n
is an integer \*(Ge 1.
.Nm
uses this special value to forward X11 connections over the secure
channel.
The user should normally not set
.Ev DISPLAY
explicitly, as that
will render the X11 connection insecure (and will require the user to
manually copy any required authorization cookies).
.It Ev HOME
Set to the path of the user's home directory.
.It Ev LOGNAME
Synonym for
.Ev USER ;
set for compatibility with systems that use this variable.
.It Ev MAIL
Set to the path of the user's mailbox.
.It Ev PATH
Set to the default
.Ev PATH ,
as specified when compiling
.Nm .
.It Ev SSH_ASKPASS
If
.Nm
needs a passphrase, it will read the passphrase from the current
terminal if it was run from a terminal.
If
.Nm
does not have a terminal associated with it but
.Ev DISPLAY
and
.Ev SSH_ASKPASS
are set, it will execute the program specified by
.Ev SSH_ASKPASS
and open an X11 window to read the passphrase.
This is particularly useful when calling
.Nm
from a
.Pa .xsession
or related script.
(Note that on some machines it
may be necessary to redirect the input from
.Pa /dev/null
to make this work.)
.It Ev SSH_AUTH_SOCK
Identifies the path of a
.Ux Ns -domain
socket used to communicate with the agent.
.It Ev SSH_CONNECTION
Identifies the client and server ends of the connection.
The variable contains
four space-separated values: client IP address, client port number,
server IP address, and server port number.
.It Ev SSH_ORIGINAL_COMMAND
This variable contains the original command line if a forced command
is executed.
It can be used to extract the original arguments.
.It Ev SSH_TTY
This is set to the name of the tty (path to the device) associated
with the current shell or command.
If the current session has no tty,
this variable is not set.
.It Ev SSH_TUNNEL
Optionally set by
.Xr sshd 8
to contain the interface names assigned if tunnel forwarding was
requested by the client.
.It Ev SSH_USER_AUTH
Optionally set by
.Xr sshd 8 ,
this variable may contain a pathname to a file that lists the authentication
methods successfully used when the session was established, including any
public keys that were used.
.It Ev TZ
This variable is set to indicate the present time zone if it
was set when the daemon was started (i.e. the daemon passes the value
on to new connections).
.It Ev USER
Set to the name of the user logging in.
.El
.Pp
Additionally,
.Nm
reads
.Pa ~/.ssh/environment ,
and adds lines of the format
.Dq VARNAME=value
to the environment if the file exists and users are allowed to
change their environment.
For more information, see the
.Cm PermitUserEnvironment
option in
.Xr sshd_config 5 .
.Sh FILES
.Bl -tag -width Ds -compact
.It Pa ~/.rhosts
This file is used for host-based authentication (see above).
On some machines this file may need to be
world-readable if the user's home directory is on an NFS partition,
because
.Xr sshd 8
reads it as root.
Additionally, this file must be owned by the user,
and must not have write permissions for anyone else.
The recommended
permission for most machines is read/write for the user, and not
accessible by others.
.Pp
.It Pa ~/.shosts
This file is used in exactly the same way as
.Pa .rhosts ,
but allows host-based authentication without permitting login with
rlogin/rsh.
.Pp
.It Pa ~/.ssh/
This directory is the default location for all user-specific configuration
and authentication information.
There is no general requirement to keep the entire contents of this directory
secret, but the recommended permissions are read/write/execute for the user,
and not accessible by others.
.Pp
.It Pa ~/.ssh/authorized_keys
Lists the public keys (DSA, ECDSA, Ed25519, RSA)
that can be used for logging in as this user.
The format of this file is described in the
.Xr sshd 8
manual page.
This file is not highly sensitive, but the recommended
permissions are read/write for the user, and not accessible by others.
.Pp
.It Pa ~/.ssh/config
This is the per-user configuration file.
The file format and configuration options are described in
.Xr ssh_config 5 .
Because of the potential for abuse, this file must have strict permissions:
read/write for the user, and not writable by others.
.Pp
.It Pa ~/.ssh/environment
Contains additional definitions for environment variables; see
.Sx ENVIRONMENT ,
above.
.Pp
.It Pa ~/.ssh/id_dsa
.It Pa ~/.ssh/id_ecdsa
.It Pa ~/.ssh/id_ed25519
.It Pa ~/.ssh/id_rsa
Contains the private key for authentication.
These files
contain sensitive data and should be readable by the user but not
accessible by others (read/write/execute).
.Nm
will simply ignore a private key file if it is accessible by others.
It is possible to specify a passphrase when
generating the key which will be used to encrypt the
sensitive part of this file using AES-128.
.Pp
.It Pa ~/.ssh/id_dsa.pub
.It Pa ~/.ssh/id_ecdsa.pub
.It Pa ~/.ssh/id_ed25519.pub
.It Pa ~/.ssh/id_rsa.pub
Contains the public key for authentication.
These files are not
sensitive and can (but need not) be readable by anyone.
.Pp
.It Pa ~/.ssh/known_hosts
Contains a list of host keys for all hosts the user has logged into
that are not already in the systemwide list of known host keys.
See
.Xr sshd 8
for further details of the format of this file.
.Pp
.It Pa ~/.ssh/rc
Commands in this file are executed by
.Nm
when the user logs in, just before the user's shell (or command) is
started.
See the
.Xr sshd 8
manual page for more information.
.Pp
.It Pa /etc/hosts.equiv
This file is for host-based authentication (see above).
It should only be writable by root.
.Pp
.It Pa /etc/shosts.equiv
This file is used in exactly the same way as
.Pa hosts.equiv ,
but allows host-based authentication without permitting login with
rlogin/rsh.
.Pp
.It Pa /etc/ssh/ssh_config
Systemwide configuration file.
The file format and configuration options are described in
.Xr ssh_config 5 .
.Pp
.It Pa /etc/ssh/ssh_host_key
.It Pa /etc/ssh/ssh_host_dsa_key
.It Pa /etc/ssh/ssh_host_ecdsa_key
.It Pa /etc/ssh/ssh_host_ed25519_key
.It Pa /etc/ssh/ssh_host_rsa_key
These files contain the private parts of the host keys
and are used for host-based authentication.
.Pp
.It Pa /etc/ssh/ssh_known_hosts
Systemwide list of known host keys.
This file should be prepared by the
system administrator to contain the public host keys of all machines in the
organization.
It should be world-readable.
See
.Xr sshd 8
for further details of the format of this file.
.Pp
.It Pa /etc/ssh/sshrc
Commands in this file are executed by
.Nm
when the user logs in, just before the user's shell (or command) is started.
See the
.Xr sshd 8
manual page for more information.
.El
.Sh EXIT STATUS
.Nm
exits with the exit status of the remote command or with 255
if an error occurred.
.Sh SEE ALSO
.Xr scp 1 ,
.Xr sftp 1 ,
.Xr ssh-add 1 ,
.Xr ssh-agent 1 ,
.Xr ssh-keygen 1 ,
.Xr ssh-keyscan 1 ,
.Xr tun 4 ,
.Xr ssh_config 5 ,
.Xr ssh-keysign 8 ,
.Xr sshd 8
.Sh STANDARDS
.Rs
.%A S. Lehtinen
.%A C. Lonvick
.%D January 2006
.%R RFC 4250
.%T The Secure Shell (SSH) Protocol Assigned Numbers
.Re
.Pp
.Rs
.%A T. Ylonen
.%A C. Lonvick
.%D January 2006
.%R RFC 4251
.%T The Secure Shell (SSH) Protocol Architecture
.Re
.Pp
.Rs
.%A T. Ylonen
.%A C. Lonvick
.%D January 2006
.%R RFC 4252
.%T The Secure Shell (SSH) Authentication Protocol
.Re
.Pp
.Rs
.%A T. Ylonen
.%A C. Lonvick
.%D January 2006
.%R RFC 4253
.%T The Secure Shell (SSH) Transport Layer Protocol
.Re
.Pp
.Rs
.%A T. Ylonen
.%A C. Lonvick
.%D January 2006
.%R RFC 4254
.%T The Secure Shell (SSH) Connection Protocol
.Re
.Pp
.Rs
.%A J. Schlyter
.%A W. Griffin
.%D January 2006
.%R RFC 4255
.%T Using DNS to Securely Publish Secure Shell (SSH) Key Fingerprints
.Re
.Pp
.Rs
.%A F. Cusack
.%A M. Forssen
.%D January 2006
.%R RFC 4256
.%T Generic Message Exchange Authentication for the Secure Shell Protocol (SSH)
.Re
.Pp
.Rs
.%A J. Galbraith
.%A P. Remaker
.%D January 2006
.%R RFC 4335
.%T The Secure Shell (SSH) Session Channel Break Extension
.Re
.Pp
.Rs
.%A M. Bellare
.%A T. Kohno
.%A C. Namprempre
.%D January 2006
.%R RFC 4344
.%T The Secure Shell (SSH) Transport Layer Encryption Modes
.Re
.Pp
.Rs
.%A B. Harris
.%D January 2006
.%R RFC 4345
.%T Improved Arcfour Modes for the Secure Shell (SSH) Transport Layer Protocol
.Re
.Pp
.Rs
.%A M. Friedl
.%A N. Provos
.%A W. Simpson
.%D March 2006
.%R RFC 4419
.%T Diffie-Hellman Group Exchange for the Secure Shell (SSH) Transport Layer Protocol
.Re
.Pp
.Rs
.%A J. Galbraith
.%A R. Thayer
.%D November 2006
.%R RFC 4716
.%T The Secure Shell (SSH) Public Key File Format
.Re
.Pp
.Rs
.%A D. Stebila
.%A J. Green
.%D December 2009
.%R RFC 5656
.%T Elliptic Curve Algorithm Integration in the Secure Shell Transport Layer
.Re
.Pp
.Rs
.%A A. Perrig
.%A D. Song
.%D 1999
.%O International Workshop on Cryptographic Techniques and E-Commerce (CrypTEC '99)
.%T Hash Visualization: a New Technique to improve Real-World Security
.Re
.Sh AUTHORS
OpenSSH is a derivative of the original and free
ssh 1.2.12 release by Tatu Ylonen.
Aaron Campbell, Bob Beck, Markus Friedl, Niels Provos,
Theo de Raadt and Dug Song
removed many bugs, re-added newer features and
created OpenSSH.
Markus Friedl contributed the support for SSH
protocol versions 1.5 and 2.0.
`