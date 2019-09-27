import 'package:flutter/material.dart';
import 'package:flutter/services.dart'; // for platform exception

/* To change the status & navigation bar color */
import 'package:flutter_statusbarcolor/flutter_statusbarcolor.dart';
import 'package:provider/provider.dart';

import 'rootPage.dart';
import 'rootModel.dart';

// Palette:
// https://www.materialpalette.com
const int PRIMARY_COLOR = 0xff010101; // Change this (format - ARGB)

void main() {
  runApp(App());
  changeNavbarColor();
}

class App extends StatelessWidget {
  final data = AppData();

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: data.title,
      theme: data.themeData,
      debugShowCheckedModeBanner: false,
      home: ChangeNotifierProvider<RootModel>(
        builder: (_) => RootModel(),
        child: RootPage(data: data),
      ),
    );
  }
}

class AppData {
  String title;
  Map<int, Color> swatchColor;
  ThemeData themeData;
  MaterialColor colorCustom;

  AppData() {
    title = 'Flutter Demo';
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
    colorCustom = MaterialColor(PRIMARY_COLOR, swatchColor);
    themeData = ThemeData(
      primaryColor: colorCustom,
      accentColor: colorCustom,
      brightness: Brightness.dark,
      scaffoldBackgroundColor: colorCustom,
      fontFamily: "OpenSans",
    );
  }
}

void changeNavbarColor() async {
  try {
    await FlutterStatusbarcolor.setNavigationBarColor(Colors.black,
        animate: true);
  } on PlatformException catch (e) {
    debugPrint(e.toString());
  }
}
