import 'package:flutter/material.dart';
import 'package:flutter/cupertino.dart';

import '../config/routes.dart' show routes;
import '../config/themes.dart' show brightAppTheme, darkAppTheme;

class MainApp extends StatelessWidget {
  const MainApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      localizationsDelegates: const <LocalizationsDelegate<dynamic>>[
        DefaultMaterialLocalizations.delegate,
        DefaultCupertinoLocalizations.delegate,
      ],
      debugShowCheckedModeBanner: false,
      theme: brightAppTheme,
      darkTheme: darkAppTheme,
      initialRoute: '/',
      routes: routes,
      themeMode: ThemeMode.system,
    );
  }
}
