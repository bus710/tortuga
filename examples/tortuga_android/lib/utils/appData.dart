/* This defines some global variables for the app's space */

import 'package:flutter/material.dart';
import 'package:flutter/services.dart'; // for platform exception

/* To change the status & navigation bar color */
import 'package:flutter_statusbarcolor/flutter_statusbarcolor.dart';

class AppData {
  // Palette: https://www.materialpalette.com
  int primaryColor;
  Map<int, Color> swatchColor;
  MaterialColor colorCustom;
  ThemeData themeData;
  String title;
  String fontFamily;

  AppData() {
    // primaryColor = 0xff00796b; // ARGB
    primaryColor = 0xff303f9f;// ARGB
    swatchColor = {
      50: Color.fromRGBO(136, 14, 79, .1),
      100: Color.fromRGBO(136, 14, 79, .2),
      200: Color.fromRGBO(136, 14, 79, .3),
      300: Color.fromRGBO(136, 14, 79, .4),
      400: Color.fromRGBO(136, 14, 79, .5),
      500: Color.fromRGBO(136, 14, 79, .6),
      600: Color.fromRGBO(136, 14, 79, .7),
      700: Color.fromRGBO(136, 14, 79, .8),
      800: Color.fromRGBO(136, 14, 79, .9),
      900: Color.fromRGBO(136, 14, 79, 1),
    };
    colorCustom = MaterialColor(primaryColor, swatchColor);
    fontFamily = "OpenSans";
    themeData = ThemeData(
      primaryColor: colorCustom,
      accentColor: colorCustom,
      brightness: Brightness.dark,
      scaffoldBackgroundColor: colorCustom,
      fontFamily: fontFamily,
    );
    title = 'Flutter Demo';
  }

  void changeNavbarColor() async {
    try {
      await FlutterStatusbarcolor.setNavigationBarColor(Color(primaryColor),
          animate: true);
    } on PlatformException catch (e) {
      debugPrint(e.toString());
    }
  }
}
