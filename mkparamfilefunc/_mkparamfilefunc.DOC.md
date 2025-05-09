<!-- Created by mkdoc DO NOT EDIT. -->

# mkparamfilefunc

This creates a Go file defining functions which set the default parameter files
for the package or program\. These can be passed as another argument to the call
where you create the parameter set or called directly, passing the parameter set
and checking for errors\. The paths of the files are derived from the XDG config
directories and from the import path of the package\.

If a group name is given the output filename and the function names will be
derived from the group name\.

It may be called multiple times in the same package directory with different
group names and with none and each time it will generate the appropriate files,
overwriting any previous files with the same name



<!-- This file is inserted into markdown files generated by mkdoc -->
<!-- if the program being documented depends on this module       -->
<!-- ============================================================ -->
<!-- See github.com/snivelingsa/utilities/mkdoc                     -->
## Parameters

This uses the `param` package and so it has access to the help parameters
which give a comprehensive message describing the usage of the program and
the parameters you can give. The `-help` parameter on its own will print the
standard parameters that the program can accept but you can also give
parameters to show both more or less help, in more or less detail. Other
standard parameters allow you to explore where parameters have been set and
where they can be set. The description of the `-help` parameter is a good
place to start to explore the help available.

The intention of the `param` package is to provide complete documentation
for the program from the command line.


<!-- This file is inserted into markdown files generated by mkdoc -->
<!-- if the program being documented depends on this module       -->
<!-- ============================================================ -->
<!-- See github.com/snivelingsa/utilities/mkdoc                     -->
## Version parameters

This offers version-querying parameters that allow the user to discover the
program version.
