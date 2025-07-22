import 'dart:convert';

import 'package:http/http.dart' as http;

import 'package:gym_genius/core/network/dio_service.dart';

import '/core/data/models/workout_dto.dart';

abstract class RemoteWorkoutDatasource {
  Future<void> saveWorkout(WorkoutDTO workout);
  Future<String> getAIDescription(WorkoutDTO workout);
}

class APIWorkoutDatasource implements RemoteWorkoutDatasource {
  APIWorkoutDatasource(this.client);
  final DioService client;

  @override
  Future<String> getAIDescription(WorkoutDTO workout) async {
    final data = {
      'model': 'cognitivecomputations/dolphin-mistral-24b-venice-edition:free',
      'messages': [
        {
          'role': 'user',
          'content': 'This is my workout: ${workout.toMap()}.'
              'Return a brief review of the workout, what should be improved and what is already good.'
              'Response format (strictly in JSON, without markdown):'
              '{“review”: “your review of the workout”}'
              'Make 3 paragraphs of text: What is good, What is bad, Summary.'
              'Add newLines between paragraphs.',
        }
      ]
    };

    final response = await http.post(
        Uri.parse('https://openrouter.ai/api/v1/chat/completions'),
        body: jsonEncode(data),
        headers: {
          'Authorization':
              'Bearer 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',
        });

    return response.body;
  }

  @override
  Future<void> saveWorkout(WorkoutDTO workout) async {
    client.post('/workouts', data: workout.toJson());
  }
}

// static Future<String> getAnalysis(
//       List<Workout> workouts, String userGoal) async {
//     try {
//       final response = await http.post(
//           Uri.parse(
//             "https://openrouter.ai/api/v1/chat/completions",
//           ),
//           headers: {
//             'Authorization': 'Bearer $_apiKey',
//           },
//           body: jsonEncode({
//             'model':
//                 'cognitivecomputations/dolphin-mistral-24b-venice-edition:free',
//             'messages': [
//               {
//                 'role': 'user',
//                 'content':
//                     'This is my workouts: $workouts. This is my goal: $userGoal. Give me analysis how am I doing. Provide 2-3 sentences.',
//               }
//             ]
//           }));

//       print(response.body);

//       final Map<String, dynamic> decoded = jsonDecode(response.body);

//       // OpenRouter/ChatGPT-like APIs usually return choices[0].message.content
//       final message = decoded['choices']?[0]?['message']?['content'];

//       return message ?? 'No recommendation received.';
//     } catch (e) {
//       print('Something wrong: $e');
//       return 'Something wrong: $e';
//     }
//   }
