// package codeometersys - gathers default values and useful application system functions.
// --
//                          Copyright (C) 2020 by Rafael Santiago
//
// Use of this source code is governed by GPL-v2 license that can
// be found in the COPYING file.
package codeometersys

// Application's version.
const appVersion = `1`

// Default help banner.
const helpBanner = `
                                                      ,  #
                                                      ########Y
                                                      ###########W
                                                       I##############
             ####WMV+                                     ,##########
           #############W                                t#########Y
         =#######=   B#####                         iB################
        ############    R###                       ################
      ################W   ###                R#####################
    t###################t ####                  ###########  #####t
   #######X i#################              ##############    ####
 W#######.  ##    :###########              +############R
V##. ###V  ###################          i################
###  ###   ##   W#############t     R####################t
M##I ##   Y##     ###################################  ###
,###      ###  :###Y;################################  ###
 ####R    ######    .################################  #Y
         ######B####################################
        +#######X    M############################=
        ######    ###############################
        #####     V############################i
         B###      ###########################
           +#        ########################
                   =##Mi  i##################              ,###W
                              ###############W          R#######R
                                 .#####,  ################V:;IW#R
                                  ###I     ###BIiY##  ##
                                ####                 =#
                                ####
                                 ##
                                  #
                              ##. #
                               ,###
                            II=  ##
                           X#######
                                ####V,
                                # ###
                                Y= .#;  C o d e o m e t e r  is Copyright (C) 2020 by Rafael Santiago.

Bug reports, feedbacks, ideas, etc: <https://github.com/rafael-santiago/codeometer/issues>
_____
usage: codeometer <command> [options]

* Are you a newbie? Oh! Welcome newbie! Give 'codeometer man' or 'codeometer help <command>' a try.
`

type CodeometerHandlerFunc func()int

func commands() map[string]CodeometerHandlerFunc {
    return map[string]CodeometerHandlerFunc {
        "measure" : measure,
        "httpd" : httpd,
        "man" : man,
        "help" : showHelpBanner,
        "version" : showAppVersion,
        "" : showHelpBanner,
    }
}
