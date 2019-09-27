import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

import 'main.dart';
import 'rootModel.dart';

class RootPage extends StatelessWidget {
  final AppData data;

  RootPage({Key key, this.data}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    var rootModel = Provider.of<RootModel>(context);

    switch (rootModel.getStatus()) {
      case Status.disconnected:
        return InitPage(data: data);
        break;
      case Status.connected:
        return InitPage(data: data);
        break;
      case Status.dial:
        return InitPage(data: data);
        break;
      case Status.init:
      default:
        return InitPage(data: data);
    }
  }
}

class InitPage extends StatefulWidget {
  final AppData data;

  InitPage({Key key, this.data}) : super(key: key);

  @override
  State<StatefulWidget> createState() {
    return InitPageState();
  }
}

class InitPageState extends State<InitPage> {
  @override
  void initState() {
    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    var rootModel = Provider.of<RootModel>(context);
    return Scaffold(
      // Give 0 to remove appBar from the actual screen
      appBar: PreferredSize(
        preferredSize: Size.fromHeight(0),
        child: AppBar(
          title: Text(widget.data.title,
              style: TextStyle(fontWeight: FontWeight.w100, fontSize: 12)),
          centerTitle: true,
          elevation: 0,
        ),
      ),
      body: Center(
        child: Container(
          decoration: BoxDecoration(
            color: Theme.of(context).primaryColor,
          ),
          child: Column(
            mainAxisAlignment: MainAxisAlignment.center,
            crossAxisAlignment: CrossAxisAlignment.center,
            children: <Widget>[
              Container(
                margin: EdgeInsets.all(10),
                child: SizedBox(
                  width: 150,
                  height: 60,
                  child: Text(
                    "Host IP",
                    style: TextStyle(fontSize: 32),
                    textAlign: TextAlign.center,
                  ),
                ),
              ),
              Container(
                width: 300,
                height: 64,
                margin: EdgeInsets.all(10),
                child: TextFormField(
                  textAlign: TextAlign.center,
                  keyboardType: TextInputType.number,
                  initialValue: "",
                  style: TextStyle(
                      fontSize: 24,
                      color: Colors.black,
                      fontFamily: "OpenSans"),
                  strutStyle: StrutStyle.fromTextStyle(TextStyle(height: 0.1)),
                  decoration: InputDecoration(
                    filled: true,
                    fillColor: Colors.grey[100],
                    border: OutlineInputBorder(
                        borderSide: BorderSide(color: Colors.white),
                        borderRadius: BorderRadius.circular(5)),
                    focusedBorder: OutlineInputBorder(
                        borderRadius: BorderRadius.circular(5)),
                    enabledBorder: OutlineInputBorder(
                        borderRadius: BorderRadius.circular(5)),
                  ),
                  onFieldSubmitted: (v) {
                    debugPrint(">>> " + v);
                    debugPrint("${rootModel.getCounter().toString()}");
                    rootModel.pressHandler();
                  },
                ),
              ),
              Container(
                height: 60,
              )
            ],
          ),
        ),
      ),
    );
  }
}
