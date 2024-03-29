===========
ego profile
===========

---------------------------------------------
LiGurOS Linux Ego Profile Module
---------------------------------------------

:Author: Daniel Robbins <drobbins@funtoo.org>
:Version: ##VERSION##
:Manual section: 1
:Manual group: LiGurOS Linux Core System

SYNOPSIS
========

The ``ego profile`` or ``epro`` command is the official way to manage LiGurOS Linux profiles.

``ego profile [command]``

``epro [command]``

``epro show``

``epro list``

``epro list [flavor|mix-ins|subarch|build]``

``epro flavor +flavor1 -flavor2..``

``epro mix-in[s] +mix-in1 -mix-in2...``

HISTORY
=======

Historically, users have had to add a large number of settings to ``/etc/make.conf`` to customize their Gentoo or LiGurOS
Linux system, which made setup of the operating system more difficult than it should be. In Gentoo Linux, it is possible
to only define one system profile. Think of a system profile as the default settings that Portage uses for building
everything on your system. Rather than a single profile, LiGurOS Linux uses multiple sub-profiles, and moves many types
of settings that are normally stored in ``/etc/make.conf`` into the profile system.

PROFILE TYPES
=============

The following profile types are available in LiGurOS Linux:

**arch**
  Typically x86-32bit or x86-64bit, this defines the processor type and support of your system. This is defined when
  your stage was built and should not be changed.

**build**
  Defines whether your system is a current, stable or experimental build. current systems will have newer packages
  unmasked than stable systems. This is defined when your stage is built and is typically not changed. Note that
  currently, only ``liguros-current`` is being used.

**subarch**
  Defines CPU optimizations for your system. The subarch is set at the time the stage3 is built, but can be changed
  later to better settings if necessary. Be sure to pick a setting that is compatible with your CPU.

**flavor**
  Defines the general type of system, such as server or desktop, and will set default USE flags appropriate for your
  needs.

**mix-ins**
 Defines various optional settings that you may be interested in enabling.

Note that flavors can (and often do) inherit settings from other flavors and mix-ins. Mix-ins can also inherit settings
from other mix-ins. For example, the ``X`` mix-in is inherited by ``gnome``. The ``epro`` tool will show both as being
enabled so there are no surprises.

USING EPRO
==========

``epro show`` will display the current profile settings on your system::

 # epro show

 === Enabled Profiles: ===
        arch: x86-64bit
       build: current
     subarch: intel64-westmere
      flavor: desktop
     mix-ins: gnome


 === Python kit: ===
      branch: 3.4-prime

 === All inherited flavors from desktop flavor: ===
                     workstation (from desktop flavor)
                            core (from workstation flavor)
                         minimal (from core flavor)

 === All inherited mix-ins from desktop flavor: ===
                               X (from workstation flavor)
                           audio (from workstation flavor)
                             dvd (from workstation flavor)
                           media (from workstation flavor)
      mediadevice-audio-consumer (from media mix-in)
                mediadevice-base (from mediadevice-audio-consumer mix-in)
      mediadevice-video-consumer (from media mix-in)
                mediadevice-base (from mediadevice-video-consumer mix-in)
        mediaformat-audio-common (from media mix-in)
          mediaformat-gfx-common (from media mix-in)
        mediaformat-video-common (from media mix-in)
                  console-extras (from workstation flavor)
                           print (from desktop flavor)

To view all available sub-profile settings, use ``epro list``. Enabled profiles will be highlighted in cyan. Directly
enabled profiles will be in bold and have an asterisk ``*`` appended. Sub-profiles enabled via inheritance will be
highlighted.

USAGE EXAMPLES
==============

``epro mix-ins +gnome +kde``
  Add the ``kde`` and ``gnome`` mix-ins.

``epro mix-ins -gnome``
  Remove the ``gnome`` mix-in.

``epro arch x86-64bit``
  Change the arch profile to ``x86-64bit``. You typically would not do this.

``epro subarch generic_64``
  Enable the ``generic_64`` subarch and associated compiler settings.

``epro flavor desktop``
  Change the system flavor to ``desktop``

``epro show-json``
  Output profile settings in standard JSON format, suitable for embedding.

``epro get flavor``
  Show current setting for flavor in plain-text format, suitable for scripting.

USING PROFILES IN YOUR OWN REPOSITORIES
=======================================

Ego now supports profiles that are defined in your own repositories. When using the ``epro list`` command, profiles
from overlays will be underlined and prefixed with ``overlayname:``.

``epro mix-ins +my_repo:foobar``
  Add the ``foobar`` mix-in defined in ``my_repo``.

DEFINING YOUR OWN PROFILES IN OVERLAYS
======================================

Creating your own subarches, mix-ins and flavors is now possible and easy to do in an existing overlay. First, add
a file ``profiles/ego.profiles.desc`` which should include JSON similar to the following. You can use this exact
configuration if you like::

 {
     "mix-ins" : "liguros/1.0/linux-gnu/mix-ins",
     "flavor" : "liguros/1.0/linux-gnu/flavor",
     "arch" : "liguros/1.0/linux-gnu/arch",
     "build" : "liguros/1.0/linux-gnu/build"
 }

This file defines locations for mix-ins, flavors, subarches and builds of LiGurOS Linux in your overlay. Note that
currently, it is not possible to define arches in your overlay, just subarches, but it is still necessary to define
an "arch" path in the JSON as this is used as a starting point to find subarch profiles.

Using the above config, you are now ready to create your own profile settings. You would then place your mix-ins in::

 <repo_path>/profiles/liguros/1.0/linux-gnu/mix-ins/my_custom_mixin

Or you could place new flavors in::

 <repo_path>/profiles/liguros/1.0/linux-gnu/flavor/my_new_flavor

When adding subarches, you will want to use the following path format::

 <repo_path>/profiles/liguros/1.0/linux-gnu/arch/<arch>/subarch/<subarch_profile_name>

For example, you might create this subarch for a new 64-bit AMD processor::

 <repo_path>/profiles/liguros/1.0/linux-gnu/arch/x86-64bit/subarch/ryzen4000

If you prefer not to use the ``liguros/1.0/linux-gnu`` prefix, you can adjust the ``profiles/ego.profiles.desc`` as
desired, and then you would modify the locations of your profiles accordingly.

Remember, as long as you have an entry in ``/etc/portage/repos.conf/<repo_name>`` for your repository, ego will be able
to see your custom profiles and they will show up in the ``epro list`` or ``ego profile list`` and will be underlined
for easy identification. Also please note that you should *not* use the ``ego-`` prefix when creating a repository file
in ``/etc/portage/repos.conf/``, as ``ego-``-prefixed repository files are managed by ego directly.


ADDITIONAL DOCUMENTATION
========================

Please see http://www.funtoo.org/LiGurOS_Profiles (``edoc "LiGurOS Profiles" | less``) for more detailed documentation,
including a list of all flavors, mix-ins, detailed documentation on LiGurOS Linux media mix-ins, how profile settings are
stored in LiGurOS Linux, as well as information about the history of the profile system, originally envisioned by Daniel
Robbins and brought to life by Seemant Kulleen.
