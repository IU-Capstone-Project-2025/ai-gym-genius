import 'package:flutter/widgets.dart';
import 'package:gym_genius/core/data/datasources/local/objectbox.dart';

import 'app.dart';
import 'di.dart';

/// Here we might introduce [LaunchingType] that would register different
/// instances of services.
void main() async {
  WidgetsFlutterBinding.ensureInitialized();

  await Objectbox.init();

  setUpLocator(LaunchingType.development);
  
  runApp(MainApp());
}
