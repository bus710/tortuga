/* In this page, it renders pages depend on the state of the rootModel. */

import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

import 'utils/appData.dart';
import 'utils/enums.dart';
import 'pages/connectedPage.dart';
import 'pages/dialPage.dart';
import 'pages/initPage.dart';

import 'rootModel.dart';

class RootPage extends StatelessWidget {
  final AppData data;

  RootPage({Key key, this.data}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    var rootModel = Provider.of<RootModel>(context);

    switch (rootModel.getStatus()) {
      case Status.connected:
        return ConnectedPage(data: data);
      case Status.connecting:
        return DialPage(data: data);
      case Status.init:
      case Status.disconnected:
      default:
        return InitPage(data: data);
    }
  }
}