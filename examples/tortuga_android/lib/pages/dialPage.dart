/* Because of the inital and loaded location of the widgets, 
  there can be a glitch of this page. To avoid that, 
  this page should have a short delay (100 mS as the timer param)
  and show the actual widgets.
  When the _show is FALSE, the page doesn't show anything.
  Timer fires after 100 mS and _show will be TRUE to print out the widgets */

import 'dart:async';

import 'package:flutter/material.dart';
import 'package:flutter_spinkit/flutter_spinkit.dart';
import 'package:provider/provider.dart';

import '../utils/appData.dart';
import '../rootModel.dart';

class DialPage extends StatefulWidget {
  final AppData data;

  DialPage({Key key, this.data}) : super(key: key);

  @override
  State<StatefulWidget> createState() {
    return DialPageState();
  }
}

class DialPageState extends State<DialPage> {
  Timer _timer;
  bool _show;

  @override
  void initState() {
    _timer = Timer.periodic(Duration(milliseconds: 100), timerHandler);
    _show = false;
    super.initState();
  }

  @override
  void dispose() {
    _timer.cancel();
    super.dispose();
  }

  void timerHandler(Timer timer) async {
    _show = true;
    setState(() {});
  }

  @override
  Widget build(BuildContext context) {
    var rootModel = Provider.of<RootModel>(context);

    if (_show) {
      return Scaffold(
        appBar: PreferredSize(
          preferredSize: Size.fromHeight(0),
          child: AppBar(),
        ),
        body: Center(
          child: Column(
            mainAxisAlignment: MainAxisAlignment.center,
            children: <Widget>[
              Container(
                height: 60,
              ),
              Container(
                decoration: BoxDecoration(
                  color: Theme.of(context).primaryColor,
                ),
                child: SpinKitFoldingCube(
                  color: Colors.grey[300],
                  size: 80.0,
                ),
              ),
              Container(
                height: 100,
              ),
            ],
          ),
        ),
      );
    } else {
      return Scaffold(
        appBar: PreferredSize(
          preferredSize: Size.fromHeight(0),
          child: AppBar(),
        ),
        body: Center(
          child: Container(
            decoration: BoxDecoration(
              color: Theme.of(context).primaryColor,
            ),
            child: Center(),
          ),
        ),
      );
    }
  }
}
