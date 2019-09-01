import 'package:flutter_web/material.dart';

void main() => runApp(MyApp());

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Tortuga Demo',
      theme: ThemeData(
        primarySwatch: Colors.blue,
        fontFamily: 'OpenSans',
      ),
      home: MyHomePage(title: 'Tortuga'),
      debugShowCheckedModeBanner: false,
    );
  }
}

class MyHomePage extends StatelessWidget {
  MyHomePage({Key key, this.title}) : super(key: key);

  final String title;

  double maxW;
  double maxH;

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text(title),
        centerTitle: true,
      ),
      body: Center(
        child: LayoutBuilder(builder: (context, constraint) {
          maxW = constraint.maxWidth;
          maxH = constraint.maxHeight;
          print(maxW.toString() + " / " + maxH.toString());
          return _getColumn();
        }),
      ),
    );
  }

  Widget _getColumn() {
    return Column(
      mainAxisAlignment: MainAxisAlignment.center,
      crossAxisAlignment: CrossAxisAlignment.center,
      children: <Widget>[
        Container(
          width: maxW - 10,
          height: maxH / 3,
          child: Center(
            child: Text(
              'Hello, World!',
            ),
          ),
        ),
        Container(
          width: maxW - 10,
          height: maxW - 10,
          decoration: BoxDecoration(
            color: Colors.white,
            border: Border.all(color: Colors.black, width: 1),
            borderRadius: BorderRadius.all(Radius.circular(3)),
          ),
        ),
      ],
    );
  }
}
