/* This is the initial page that accepts an IP address of the robot. */

import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

import '../utils/appData.dart';
import '../utils/enums.dart';
import '../rootModel.dart';

class InitPage extends StatefulWidget {
  final AppData data;

  InitPage({Key key, this.data}) : super(key: key);

  @override
  State<StatefulWidget> createState() {
    return InitPageState();
  }
}

class InitPageState extends State<InitPage> {
  String mainText;

  @override
  void initState() {
    mainText = "Host IP";
    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    var rootModel = Provider.of<RootModel>(context);

    return Scaffold(
      // Give 0 to remove appBar from the actual screen
      appBar: PreferredSize(
        preferredSize: Size.fromHeight(0),
        child: AppBar(),
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
                margin: EdgeInsets.all(40),
                child: SizedBox(
                  width: 150,
                  height: 60,
                  child: Text(
                    mainText,
                    style: TextStyle(
                        fontSize: 44,
                        fontWeight: FontWeight.w100,
                        color: Colors.grey[100]),
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
                      fontFamily: widget.data.fontFamily),
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
                  onFieldSubmitted: (address) {
                    rootModel.pressHandler(Request.dial, address);
                  },
                ),
              ),
              Container(
                height: 100,
              )
            ],
          ),
        ),
      ),
    );
  }
}
