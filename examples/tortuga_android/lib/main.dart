import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

import 'utils/appData.dart';
import 'rootPage.dart';
import 'rootModel.dart';

void main() {
  runApp(App());
}

class App extends StatelessWidget {
  final data = AppData();

  @override
  Widget build(BuildContext context) {
    data.changeNavbarColor();
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

