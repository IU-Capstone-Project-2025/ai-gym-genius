import 'package:flutter/material.dart';

final ThemeData darkAppTheme = ThemeData(
  colorScheme: ColorScheme(
    primary: Color(0xFFFF6E00), // Bright orange
    primaryContainer: Colors.amber, // Energetic red-orange
    secondary: const Color(0xFF121212), // Dark background
    secondaryContainer: Color(0xFF2A2A2A), // Dark variant
    surface: Color(0xFF1E1E1E), // Background
    error: Color(0xFFD32F2F), // Error red
    onPrimary: Colors.black, // Text/icons on primary
    onSecondary: Colors.white, // Text/icons on secondary
    onSurface: Colors.white, // Text/icons on surface
    onError: Colors.white, // Text/icons on error
    brightness: Brightness.dark, // Dark theme base
  ),
);

final ThemeData brightAppTheme = ThemeData(
  colorScheme: ColorScheme(
    primary: Color(0xFFFF6E00), // Bright orange
    primaryContainer: Colors.amber, // Energetic red-orange
    secondary: Colors.white, // Dark background
    secondaryContainer: Colors.grey.shade200,
    surface: Colors.white, // Background
    error: Color(0xFFD32F2F), // Error red
    onPrimary: Colors.black, // Text/icons on primary
    onSecondary: Colors.black, // Text/icons on secondary
    onSurface: Colors.black, // Text/icons on surface
    onError: Colors.black, // Text/icons on error
    brightness: Brightness.light, // Dark theme base
  ),
);
