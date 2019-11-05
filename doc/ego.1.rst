=========
ego
=========

---------------------------------------------
LiGurOS Linux System Management Meta-Command
---------------------------------------------

:Author: Daniel Robbins <drobbins@funtoo.org>
:Version: ##VERSION##
:Manual section: 1
:Manual group: LiGurOS Linux Core System

SYNOPSIS
--------

The *ego* command is a meta-command that consists of several modules that provide functionality. Currently, the
following modules are supported:

sync
  Ego sync module. See ego-sync(8).

profile
  Ego profile querying and selection module (shortcut: *epro*) See ego-profile(8).

query
  Ego query module. See ego-query(1).

doc
  Ego Wiki Documentation module. (shortcut: *edoc*) See ego-doc(1).

kit
  Ego kit information module. see ego-kit(1).

config
  Ego configuration module to allow changing ``/etc/ego.conf`` from the command-line. See ego-config(8).

You can invoke the relevant module by using the calling convention ``ego module [arg1...]``. For example, to view the
LiGurOS Linux Installation Guide, type ``ego doc Install | less``. Alternatively, you can use the module shortcut if
one exists, such as ``edoc install | less``.

ENVIRONMENT VARIABLES
---------------------

``EGO_CONFIG``
  Use this environment variable to set an alternate path to the configuration file (default is ``/etc/ego.conf``.)

``ROOT``
  Use this environment variable to set an alternate path for the OS installation. This allows you to use your local
  ego command to interact with a LiGurOS system in a chroot, for example. Note that when this environment variable is
  set, ego will not drop permissions to the Portage user.

LICENSE
--------

.. include:: ../COPYRIGHT.txt
