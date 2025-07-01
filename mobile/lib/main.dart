import 'package:flutter/widgets.dart';

import 'app.dart';
import 'di.dart';

/// Here we might introduce [LaunchingType] that would register different
/// instances of services.
void main() async {
  // I introduce dev
  setUpLocator(LaunchingType.development);
  runApp(MainApp());
}
