import 'package:flutter/material.dart';

class ActivityGrid extends StatefulWidget {
  final Set<DateTime> workoutDays; // normalised to Y-M-D
  final double cellSize;
  final double spacing;

  const ActivityGrid({
    super.key,
    required this.workoutDays,
    this.cellSize = 14,
    this.spacing = 6,
  });

  @override
  State<ActivityGrid> createState() => _ActivityGridState();
}

class _ActivityGridState extends State<ActivityGrid> {
  late final ScrollController _scrollCtrl;
  late final double gridHeight;

  @override
  void initState() {
    super.initState();
    gridHeight = widget.cellSize * 7 + widget.spacing * 20;
    _scrollCtrl = ScrollController();

    // Scroll to the newest week after first layout
    WidgetsBinding.instance.addPostFrameCallback((_) {
      if (!_scrollCtrl.hasClients) return;

      final max = _scrollCtrl.position.maxScrollExtent;

      _scrollCtrl.animateTo(
        max, // top in reversed axis
        duration: const Duration(milliseconds: 1200),
        curve: Curves.decelerate,
      );
    });
  }

  @override
  void dispose() {
    _scrollCtrl.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    final today = DateTime.now();
    final firstDay = today.subtract(const Duration(days: 7 * 52 - 1));
    final schema = Theme.of(context).colorScheme;

    // Build a list of every day from firstDay → today
    final days = List<DateTime>.generate(
      7 * 52,
      (i) => DateTime(firstDay.year, firstDay.month, firstDay.day + i),
    );

    return Container(
      padding: EdgeInsets.only(left: 32, right: 8, top: 24),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          Text(
            'This is your activity github (joke)',
            style: Theme.of(context).textTheme.titleMedium,
          ),
          SizedBox(
            height: 8,
          ),
          IntrinsicHeight(
            child: Row(
              children: [
                Expanded(
                  flex: 9,
                  child: SizedBox(
                    height: gridHeight,
                    child: GridView.builder(
                      controller: _scrollCtrl,
                      scrollDirection: Axis.horizontal,
                      reverse: false,
                      shrinkWrap: true,
                      gridDelegate: SliverGridDelegateWithFixedCrossAxisCount(
                        crossAxisCount: 7, // columns = 7 days (Mon–Sun)
                        mainAxisSpacing: widget.spacing,
                        crossAxisSpacing: widget.spacing,
                      ),
                      itemCount: days.length,
                      itemBuilder: (_, i) {
                        final d = days[i];
                        final bool done =
                            widget.workoutDays.any((w) => isSameDay(w, d));
            
                        return Container(
                          width: widget.cellSize,
                          height: widget.cellSize,
                          decoration: BoxDecoration(
                            color:
                                done ? schema.primary : schema.secondaryContainer,
                            borderRadius: BorderRadius.circular(3),
                          ),
                        );
                      },
                    ),
                  ),
                ),
                SizedBox(width: 16,),
                Expanded(
                  child: Column(
                    mainAxisAlignment: MainAxisAlignment.spaceBetween,
                    crossAxisAlignment: CrossAxisAlignment.start,
                    children: [
                      Text('mon'),
                      Text('tue'),
                      Text('wed'),
                      Text('thu'),
                      Text('fri'),
                      Text('sat'),
                      Text('sun'),
                    ],
                  ),
                )
              ],
            ),
          ),
        ],
      ),
    );
  }
}

bool isSameDay(DateTime a, DateTime b) =>
    a.year == b.year && a.month == b.month && a.day == b.day;
