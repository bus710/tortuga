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

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text(title),
        centerTitle: true,
      ),
      body: AppPage(title: this.title),
    );
  }
}

class AppPage extends StatefulWidget {
  AppPage({Key key, this.title}) : super(key: key);

  final String title;

  @override
  State<StatefulWidget> createState() {
    return _AppState();
  }
}

class _AppState extends State<AppPage> {
  double maxW;
  double maxH;

  double padW = 0;
  double widX = 0;
  double widY = 0;

  double x = 0;
  double y = 0;

  @override
  Widget build(BuildContext context) {
    return Center(
      child: LayoutBuilder(builder: (context, constraint) {
        maxW = constraint.maxWidth;
        maxH = constraint.maxHeight;
        // print(maxW.toString() + " / " + maxH.toString());
        return _getColumn();
      }),
    );
  }

  Widget _getColumn() {
    if (maxW > 512) {
      padW = 512;
    } else {
      padW = maxW;
    }

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
          width: padW - 10,
          height: padW - 10,
          decoration: BoxDecoration(
            color: Colors.white,
            border: Border.all(color: Colors.black, width: 1),
            borderRadius: BorderRadius.all(Radius.circular(3)),
          ),
          child: Stack(
            children: <Widget>[
              Positioned(
                right: widX,
                top: widY,
                child: FlatButton(
                  child: Text("test"),
                  onPressed: () => {print("pressed")},
                ),
              ),
              GestureDetector(
                onPanStart: (d) => {
                  print("location: " + widX.toString() + "/" + widY.toString()),
                  x = d.localPosition.dx,
                  y = d.localPosition.dy,
                  widX = x * -1 + padW - 30,
                  widY = y,
                  print("start: " + x.toString() + "/" + y.toString()),
                  setState(() {}),
                },
                onPanUpdate: (d) => {
                  print("location: " + widX.toString() + "/" + widY.toString()),
                  x = d.localPosition.dx,
                  y = d.localPosition.dy,
                  widX = x * -1 + padW - 30,
                  widY = y,
                  print("update: " + x.toString() + "/" + y.toString()),
                  setState(() {}),
                },
                // child:
              ),
            ],
          ),
        ),
      ],
    );
  }
}
