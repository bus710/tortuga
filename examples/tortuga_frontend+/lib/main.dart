import 'package:flutter/material.dart';
import 'package:tortuga_frontend/AppPage.dart';

void main() => runApp(MyApp());

// Palette:
// https://www.materialpalette.com
const int PRIMARY_COLOR = 0xff00796b; // Change this

class MyApp extends StatelessWidget {
  // How to make a custom material color:
  // https://medium.com/@manojvirat457/turn-any-color-to-material-color-for-flutter-d8e8e037a837

  final Map<int, Color> swatchColor = {
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

  @override
  Widget build(BuildContext context) {
    MaterialColor colorCustom = MaterialColor(PRIMARY_COLOR, swatchColor);

    return MaterialApp(
      title: 'Tortuga',
      theme: ThemeData(
        primarySwatch: colorCustom,
        fontFamily: 'OpenSans',
      ),
      home: MyHomePage(title: ''),
      debugShowCheckedModeBanner: false,
    );
  }
}

class MyHomePage extends StatelessWidget {
  MyHomePage({Key key, this.title}) : super(key: key);

  final String title;

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      // Give 1 to remove appBar from the actual screen
      appBar: PreferredSize(
        preferredSize: Size.fromHeight(1),
        child: AppBar(
          title: Text(title,
              style: TextStyle(fontWeight: FontWeight.w100, fontSize: 12)),
          centerTitle: true,
          elevation: 0.0,
        ),
      ),
      body: AppPage(
        title: this.title,
      ),
    );
  }
}
