import 'package:dio/dio.dart';
import 'package:flutter/foundation.dart';

class DioService {
  // Private constructor for singleton pattern
  DioService._privateConstructor() {
    _dio = Dio(BaseOptions(
        baseUrl: "http://api.xn--b1ab5acc.site",
        connectTimeout: const Duration(seconds: 10),
        receiveTimeout: const Duration(seconds: 30),
        headers: {
          'Content-Type': 'application/json',
          'Authorization': 'Bearer 123123123'
        }));

    if (kDebugMode) {
      _dio.interceptors.add(LogInterceptor());
    }
  }

  static final DioService _instance = DioService._privateConstructor();

  factory DioService() {
    return _instance;
  }

  late final Dio _dio;

  /// Performs a GET request.
  Future<dynamic> get(String endpoint,
      {Map<String, dynamic>? queryParameters}) async {
    try {
      final response =
          await _dio.get(endpoint, queryParameters: queryParameters);
      return response.data;
    } on DioException catch (e) {
      throw _handleError(e);
    }
  }

  /// Performs a POST request.
  Future<dynamic> post(String endpoint, {dynamic data}) async {
    try {
      final response = await _dio.post(endpoint, data: data);
      return response.data;
    } on DioException catch (e) {
      throw _handleError(e);
    }
  }

  /// Performs a PUT request.
  Future<dynamic> put(String endpoint, {dynamic data}) async {
    try {
      final response = await _dio.put(endpoint, data: data);
      return response.data;
    } on DioException catch (e) {
      throw _handleError(e);
    }
  }

  /// Performs a DELETE request.
  Future<dynamic> delete(String endpoint) async {
    try {
      final response = await _dio.delete(endpoint);
      return response.data;
    } on DioException catch (e) {
      throw _handleError(e);
    }
  }
}

/// Handles Dio errors and converts them to a user-friendly [ApiException].
ApiException _handleError(DioException e) {
  String errorMessage = 'An unexpected error occurred.';

  switch (e.type) {
    case DioExceptionType.connectionTimeout:
    case DioExceptionType.sendTimeout:
    case DioExceptionType.receiveTimeout:
      errorMessage =
          'Connection timeout. Please check your internet connection.';
      break;
    case DioExceptionType.badResponse:
      final statusCode = e.response?.statusCode;
      if (statusCode != null) {
        if (statusCode >= 500) {
          errorMessage = 'Server error. Please try again later.';
        } else if (statusCode >= 400) {
          // You can parse the error response body for more specific messages
          errorMessage = e.response?.data?['message'] ??
              'Bad request. Please check your input.';
        }
      }
      break;
    case DioExceptionType.cancel:
      errorMessage = 'Request was cancelled.';
      break;
    case DioExceptionType.connectionError:
      errorMessage = 'No internet connection.';
      break;
    case DioExceptionType.unknown:
    default:
      errorMessage = 'An unknown error occurred.';
      break;
  }

  return ApiException(errorMessage);
}

/// A custom exception class for API errors.
class ApiException implements Exception {
  final String message;
  ApiException(this.message);

  @override
  String toString() => message;
}
